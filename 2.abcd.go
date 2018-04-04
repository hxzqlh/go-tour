package main

import (
	"fmt"
	"os"
)

func Cor1() string {
	ch := make(chan int)
	go func() {
		fmt.Printf("1 ")
		ch <- 1
	}()
	<-ch
	return "1 "
}

func Cor2() string {
	ch := make(chan int)
	go func() {
		fmt.Printf("2 ")
		ch <- 2
	}()
	<-ch
	return "2 "
}

func Cor3() string {
	ch := make(chan int)
	go func() {
		fmt.Printf("3 ")
		ch <- 3
	}()
	<-ch
	return "3 "
}

func Cor4() string {
	ch := make(chan int)
	go func() {
		fmt.Printf("4 ")
		ch <- 4
	}()
	<-ch
	return "4 "
}

func main() {
	text_inic()
	write_to_file('A')
	write_to_file('B')
	write_to_file('C')
	write_to_file('D')
}

//打开文件并写入数据
func write_to_file(n byte) {
	userFileStr := "A.txt"
	userFile := []byte(userFileStr)
	userFile[0] += (n - 'A')
	userFileStr = string(userFile)
	fount, err := os.OpenFile(userFileStr, os.O_WRONLY|os.O_CREATE, 0666)
	defer fount.Close()
	if err != nil {
		fmt.Printf("打开文件:%s失败\n", userFileStr)
	} else {
		fmt.Printf("打开文件:%s成功: %c: ", userFileStr, n)
		// ch := make(chan int)
		for i := 0; i < 10; i++ {
			switch n - 'A' {
			case 0:
				fount.WriteString(Cor1() + Cor2() + Cor3() + Cor4())
			case 1:
				fount.WriteString(Cor2() + Cor3() + Cor4() + Cor1())
			case 2:
				fount.WriteString(Cor3() + Cor4() + Cor1() + Cor2())
			case 3:
				fount.WriteString(Cor4() + Cor1() + Cor2() + Cor3())
			}
		}
		fmt.Println("")
	}
}

//初始化创建文件
func text_inic() {
	filenameStr := "A.txt"
	filename := []byte(filenameStr)
	for i := 0; i < 4; i++ {
		fount, err := os.Create(filenameStr)
		if err != nil {
			fmt.Printf("创建文件:%s失败\n", filenameStr)
		}
		fmt.Printf("创建文件:%s成功\n", filenameStr)
		filename[0]++
		filenameStr = string(filename)
		defer fount.Close()
	}
}
