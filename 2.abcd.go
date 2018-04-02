package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)

	go func() {
		for {
			ch1 <- 1
		}
	}()

	go func() {
		for {
			ch2 <- 2
		}
	}()

	go func() {
		for {
			ch3 <- 3
		}
	}()

	go func() {
		for {
			ch4 <- 4
		}
	}()

	var filenameA = "./desktop/A.txt"
	var filenameB = "./desktop/B.txt"
	var filenameC = "./desktop/C.txt"
	var filenameD = "./desktop/D.txt"

	var fA *os.File
	var fB *os.File
	var fC *os.File
	var fD *os.File

	var err1 error

	go func() {
		if checkFileIsExist(filenameA) { //如果文件存在
			fA, err1 = os.OpenFile(filenameA, os.O_APPEND, 0666) //打开文件
			fmt.Println("文件存在")
		} else {
			fA, err1 = os.Create(filenameA) //创建文件
			fmt.Println("文件不存在")
		}
		check(err1)
		for {
			_, err1 = io.WriteString(fA, strconv.Itoa(<-ch1)) //写入文件(字符串)
			_, err1 = io.WriteString(fA, " ")                 //写入文件(字符串)
			_, err1 = io.WriteString(fA, strconv.Itoa(<-ch2)) //写入文件(字符串)
			_, err1 = io.WriteString(fA, " ")                 //写入文件(字符串)
			_, err1 = io.WriteString(fA, strconv.Itoa(<-ch3)) //写入文件(字符串)
			_, err1 = io.WriteString(fA, " ")                 //写入文件(字符串)
			_, err1 = io.WriteString(fA, strconv.Itoa(<-ch4)) //写入文件(字符串)
			_, err1 = io.WriteString(fA, " ")                 //写入文件(字符串)
		}
		check(err1)
	}()

	go func() {
		if checkFileIsExist(filenameB) { //如果文件存在
			fB, err1 = os.OpenFile(filenameB, os.O_APPEND, 0666) //打开文件
			fmt.Println("文件存在")
		} else {
			fB, err1 = os.Create(filenameB) //创建文件
			fmt.Println("文件不存在")
		}
		check(err1)
		for {
			_, err1 = io.WriteString(fB, strconv.Itoa(<-ch2)) //写入文件(字符串)
			_, err1 = io.WriteString(fB, " ")                 //写入文件(字符串)
			_, err1 = io.WriteString(fB, strconv.Itoa(<-ch3)) //写入文件(字符串)
			_, err1 = io.WriteString(fB, " ")                 //写入文件(字符串)
			_, err1 = io.WriteString(fB, strconv.Itoa(<-ch4)) //写入文件(字符串)
			_, err1 = io.WriteString(fB, " ")                 //写入文件(字符串)
			_, err1 = io.WriteString(fB, strconv.Itoa(<-ch1)) //写入文件(字符串)
			_, err1 = io.WriteString(fB, " ")                 //写入文件(字符串)
		}
		check(err1)
	}()

	go func() {
		if checkFileIsExist(filenameC) { //如果文件存在
			fC, err1 = os.OpenFile(filenameC, os.O_APPEND, 0666) //打开文件
			fmt.Println("文件存在")
		} else {
			fC, err1 = os.Create(filenameC) //创建文件
			fmt.Println("文件不存在")
		}
		check(err1)
		for {
			_, err1 = io.WriteString(fC, strconv.Itoa(<-ch3)) //写入文件(字符串)
			_, err1 = io.WriteString(fC, " ")                 //写入文件(字符串)
			_, err1 = io.WriteString(fC, strconv.Itoa(<-ch4)) //写入文件(字符串)
			_, err1 = io.WriteString(fC, " ")                 //写入文件(字符串)
			_, err1 = io.WriteString(fC, strconv.Itoa(<-ch1)) //写入文件(字符串)
			_, err1 = io.WriteString(fC, " ")                 //写入文件(字符串)
			_, err1 = io.WriteString(fC, strconv.Itoa(<-ch2)) //写入文件(字符串)
			_, err1 = io.WriteString(fC, " ")                 //写入文件(字符串)
		}
		check(err1)
	}()

	go func() {
		if checkFileIsExist(filenameD) { //如果文件存在
			fD, err1 = os.OpenFile(filenameD, os.O_APPEND, 0666) //打开文件
			fmt.Println("文件存在")
		} else {
			fD, err1 = os.Create(filenameD) //创建文件
			fmt.Println("文件不存在")
		}
		check(err1)
		for {
			_, err1 = io.WriteString(fD, strconv.Itoa(<-ch4)) //写入文件(字符串)
			_, err1 = io.WriteString(fD, " ")                 //写入文件(字符串)
			_, err1 = io.WriteString(fD, strconv.Itoa(<-ch1)) //写入文件(字符串)
			_, err1 = io.WriteString(fD, " ")                 //写入文件(字符串)
			_, err1 = io.WriteString(fD, strconv.Itoa(<-ch2)) //写入文件(字符串)
			_, err1 = io.WriteString(fD, " ")                 //写入文件(字符串)
			_, err1 = io.WriteString(fD, strconv.Itoa(<-ch3)) //写入文件(字符串)
			_, err1 = io.WriteString(fD, " ")                 //写入文件(字符串)

		}
		check(err1)
	}()

	time.Sleep(time.Second * 3)

}
