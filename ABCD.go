package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var ch1 = make(chan int, 100)
var ch2 = make(chan int, 100)
var ch3 = make(chan int, 100)
var ch4 = make(chan int, 100)

var str [7]int

func main() {
	f1, err1 := os.OpenFile("A.txt", os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm|os.ModeTemporary)
	if err1 != nil {
		fmt.Println(err1.Error())
		return
	}
	f2, err2 := os.OpenFile("B.txt", os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm|os.ModeTemporary)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}
	f3, err3 := os.OpenFile("C.txt", os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm|os.ModeTemporary)
	if err3 != nil {
		fmt.Println(err3.Error())
		return
	}
	f4, err4 := os.OpenFile("D.txt", os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm|os.ModeTemporary)
	if err4 != nil {
		fmt.Println(err4.Error())
		return
	}

	go w1()
	go w2()
	go w3()
	go w4()
	go func() {

		for i := 0; i < 50; i++ {
			t := i % 4
			str[0] = <-ch1
			str[1] = <-ch2
			str[2] = <-ch3
			str[3] = <-ch4
			str[4] = <-ch1
			str[5] = <-ch2
			str[6] = <-ch3
			f1.Write([]byte(strconv.Itoa(str[t]) + " "))
			f2.Write([]byte(strconv.Itoa(str[t+1]) + " "))
			f3.Write([]byte(strconv.Itoa(str[t+2]) + " "))
			f4.Write([]byte(strconv.Itoa(str[t+3]) + " "))

		}
	}()
	time.Sleep(200 * time.Millisecond)

}

func w1() {
	for {
		ch1 <- 1
	}
}
func w2() {
	for {
		ch2 <- 2
	}
}
func w3() {
	for {
		ch3 <- 3
	}
}
func w4() {
	for {
		ch4 <- 4
	}
}
