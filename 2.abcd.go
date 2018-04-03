package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {

	fileA, err := os.OpenFile(
		"A.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE|os.O_APPEND,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}

	fileB, err := os.OpenFile(
		"B.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE|os.O_APPEND,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	fileC, err := os.OpenFile(
		"C.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE|os.O_APPEND,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	fileD, err := os.OpenFile(
		"D.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE|os.O_APPEND,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}

	ch1 := makeChs(4)
	ch2 := makeChs(4)
	ch3 := makeChs(4)
	ch4 := makeChs(4)

	go gen(1, ch1)
	go gen(2, ch2)
	go gen(3, ch3)
	go gen(4, ch4)

	chs1 := []chan int{ch1[0], ch2[0], ch3[0], ch4[0]}
	buf1 := new(bytes.Buffer)
	go writeFile(chs1, buf1)

	chs2 := []chan int{ch2[1], ch3[1], ch4[1], ch1[1]}
	buf2 := new(bytes.Buffer)
	go writeFile(chs2, buf2)

	chs3 := []chan int{ch3[2], ch4[2], ch1[2], ch2[2]}
	buf3 := new(bytes.Buffer)
	go writeFile(chs3, buf3)

	chs4 := []chan int{ch4[3], ch1[3], ch2[3], ch3[3]}
	buf4 := new(bytes.Buffer)
	go writeFile(chs4, buf4)

	time.Sleep(2 * time.Second)
	printData(buf1, fileA)
	printData(buf2, fileB)
	printData(buf3, fileC)
	printData(buf4, fileD)
	fmt.Println("done")
}

func makeChs(len int) (rets []chan int) {
	rets = make([]chan int, len)
	for i := 0; i < len; i++ {
		rets[i] = make(chan int)
	}
	return
}

func gen(seed int, rets []chan int) {
	for {
		for _, ch := range rets {
			ch <- seed
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func writeFile(chs []chan int, buf *bytes.Buffer) {
	for {
		for _, ch := range chs {
			data := <-ch
			buf.Write([]byte(strconv.Itoa(data) + " "))
		}
	}
}

func printData(buf *bytes.Buffer, file *os.File) {
	byteSlise := buf.Bytes()
	_, err := file.Write(byteSlise)
	if err != nil {
		log.Fatal(err)
	}
}
