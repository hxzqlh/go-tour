/*1. 有四个协程1、2、3、4, 协程1的功能就是输出1，协程2的功能就是输出2，以此类推……现在有四个文件A、B、C、D，初始都为空，编程实现让这四个文件呈现如下格式：
A：1 2 3 4 1 2...
B：2 3 4 1 2 3...
C：3 4 1 2 3 4...
D：4 1 2 3 4 1...*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var ch1, ch2, ch3, ch4 chan int

func main() {
	ch1 = make(chan int)
	ch2 = make(chan int)
	ch3 = make(chan int)
	ch4 = make(chan int)
	go Ch1()
	go Ch2()
	go Ch3()
	go Ch4()
	magicNumber('A', 10)
	magicNumber('B', 10)
	magicNumber('C', 10)
	magicNumber('D', 10)
}

func magicNumber(fliename byte, round int) {
	var temp int = int(fliename - 65)
	for i := temp; i < temp+round; i++ {
		num := i%4 + 1
		switch {
		case num == 1:
			fliewrite(string(fliename), <-ch1)
		case num == 2:
			fliewrite(string(fliename), <-ch2)
		case num == 3:
			fliewrite(string(fliename), <-ch3)
		case num == 4:
			fliewrite(string(fliename), <-ch4)
		}
	}
}
func Ch1() {
	for {
		ch1 <- 1
	}
}

func Ch2() {
	for {
		ch2 <- 2
	}
}

func Ch3() {
	for {
		ch3 <- 3
	}
}

func Ch4() {
	for {
		ch4 <- 4
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func fliewrite(fn string, num int) {
	var filename = "./" + fn + ".txt"
	var f *os.File
	var err1 error
	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}
	check(err1)
	w := bufio.NewWriter(f) //创建新的 Writer 对象
	w.WriteString(strconv.Itoa(num))
	//fmt.Printf("写入 %d 个字节n", n4)
	w.Flush()
	f.Close()
}
