package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func createFile(filename string) *os.File {
	var f *os.File
	var err error

	if checkFileIsExist(filename) { //如果文件存在
		f, err = os.OpenFile(filename, os.O_RDWR, 0666) //打开文件
		fmt.Println(filename, "文件存在")
	} else {
		f, err = os.Create(filename) //创建文件
		fmt.Println(filename, "文件不存在")
	}

	if err != nil {
		panic(err)
	}

	return f
}

func charCreator(v string, done chan bool) <-chan string {
	c := make(chan string)
	go func(s string, done chan bool) {
		for {
			select {
			case c <- s:
			case <-done:
				close(c)
				return
			}
		}
	}(v, done)

	return c
}

//SenderNum ...
type SenderNum int

const (
	//Min ...
	Min SenderNum = 0
	//One ...
	One SenderNum = 1
	//Two ...
	Two SenderNum = 2
	//Three ...
	Three SenderNum = 3
	//Four ...
	Four SenderNum = 4
	//Max ...
	Max SenderNum = 5
)

type sendBuf struct {
	f          *os.File
	nextSender SenderNum
}

func main() {
	var bufs []sendBuf
	var tmpBuf sendBuf

	//Create A file and relate buffer
	tmpBuf.f = createFile("./A")
	tmpBuf.nextSender = One
	bufs = append(bufs, tmpBuf)
	//Create B file and relate buffer
	tmpBuf.f = createFile("./B")
	tmpBuf.nextSender = Two
	bufs = append(bufs, tmpBuf)
	//Create C file and relate buffer
	tmpBuf.f = createFile("./C")
	tmpBuf.nextSender = Three
	bufs = append(bufs, tmpBuf)
	//Create D file and relate buffer
	tmpBuf.f = createFile("./D")
	tmpBuf.nextSender = Four
	bufs = append(bufs, tmpBuf)

	defer func() {
		for _, v := range bufs {
			if v.f != nil {
				v.f.Close()
			}
		}
	}()

	done := make(chan bool)
	senders := make(map[SenderNum]<-chan string)
	senders[One] = charCreator("1", done)
	senders[Two] = charCreator("2", done)
	senders[Three] = charCreator("3", done)
	senders[Four] = charCreator("4", done)

	var stop bool
	go func(done chan bool) {
		select {
		//期望数据注入的时间，当前为1s;
		case <-time.After(time.Second * 1):
			close(done)
		}
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		stop = true
	}(done)

	for !stop {
		for k, buf := range bufs {
			_, err := io.WriteString(buf.f, <-senders[buf.nextSender]) //写入文件(字符串)
			if err != nil {
				fmt.Println(err.Error())
			}
			if buf.nextSender+1 >= Max {
				bufs[k].nextSender = One
			} else {
				bufs[k].nextSender = buf.nextSender + 1
			}
		}
	}
	fmt.Println("Stoped!")
	return
}


