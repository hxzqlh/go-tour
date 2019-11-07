package main

import (
	"io"
	"os"
	"strconv"
	"sync"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

var wg1 = sync.WaitGroup{}

func writeNum(data *MyData, num int) {
	defer wg1.Done()
	v := num + 1
	if num == 4 {
		v = 1
	}
	for i := 0; i < 100000; i++ {
		if x, ok := <-data.ch1; ok {
			if num == x {
				io.WriteString(data.f1, strconv.Itoa(num))
				data.ch1 <- v
			} else {
				data.ch1 <- x
			}
		}
		if x, ok := <-data.ch2; ok {
			if num == x {
				io.WriteString(data.f2, strconv.Itoa(num))
				data.ch2 <- v
			} else {
				data.ch2 <- x
			}
		}
		if x, ok := <-data.ch3; ok {
			if num == x {
				io.WriteString(data.f3, strconv.Itoa(num))
				data.ch3 <- v
			} else {
				data.ch3 <- x
			}
		}
		if x, ok := <-data.ch4; ok {
			if num == x {
				io.WriteString(data.f4, strconv.Itoa(num))
				data.ch4 <- v
			} else {
				data.ch4 <- x
			}
		}
	}

}

type MyData struct {
	f1, f2, f3, f4     *os.File
	ch1, ch2, ch3, ch4 chan int
}

//整体思路为  创建4个通道分别对应四个文件下次改写入那个数字。
//				然后每个写数字的goroutine循环读取每个通道，
//				如果该通道的值和自己负责写的值一致，就写入通道对应的文件
//					并将下一个该写入的值存入通道
//				否则如果不一致，则把从通道读出来的值继续存入通道
func main() {
	//创建/清空文件
	_, err := os.Create("./A.txt")
	_, err = os.Create("./B.txt")
	_, err = os.Create("./C.txt")
	_, err = os.Create("./D.txt")
	checkErr(err)

	//打开文件
	f1, err := os.OpenFile("./A.txt", os.O_APPEND, 0666)
	f2, err := os.OpenFile("./B.txt", os.O_APPEND, 0666)
	f3, err := os.OpenFile("./C.txt", os.O_APPEND, 0666)
	f4, err := os.OpenFile("./D.txt", os.O_APPEND, 0666)
	defer f1.Close()
	defer f2.Close()
	defer f3.Close()
	defer f4.Close()
	//数据
	myData := MyData{
		f1:  f1,
		f2:  f2,
		f3:  f3,
		f4:  f4,
		ch1: make(chan int, 1),
		ch2: make(chan int, 1),
		ch3: make(chan int, 1),
		ch4: make(chan int, 1),
	}
	//初始操作把第一次内容写入并把下一次要写的内容放入通道
	io.WriteString(myData.f1, "1")
	myData.ch1 <- 2
	io.WriteString(myData.f2, "2")
	myData.ch2 <- 3
	io.WriteString(myData.f3, "3")
	myData.ch3 <- 4
	io.WriteString(myData.f4, "4")
	myData.ch4 <- 1

	wg1.Add(4)
	go writeNum(&myData, 1)
	go writeNum(&myData, 2)
	go writeNum(&myData, 3)
	go writeNum(&myData, 4)
	wg1.Wait()
}
