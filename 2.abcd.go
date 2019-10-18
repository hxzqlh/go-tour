package main

import (
	"io/ioutil"
)

func out1(ch chan string) {
	for {
		ch <- "1"
	}
}

func out2(ch chan string) {
	for {
		ch <- "2"
	}
}

func out3(ch chan string) {
	for {
		ch <- "3"
	}
}

func out4(ch chan string) {
	for {
		ch <- "4"
	}
}

func wrt(fn string, quit chan byte, ch1, ch2, ch3, ch4 chan string) {
	var sUnit string
	for i := 0; i < 10; i++ {
		sUnit += <-ch1 + <-ch2 + <-ch3 + <-ch4
	}
	ioutil.WriteFile(fn, []byte(sUnit), 0755)
	quit <- byte(0)
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	ch4 := make(chan string)
	quit := make(chan byte)
	go out1(ch1)
	go out2(ch2)
	go out3(ch3)
	go out4(ch4)

	go wrt("./A.txt", quit, ch1, ch2, ch3, ch4)
	go wrt("./B.txt", quit, ch2, ch3, ch4, ch1)
	go wrt("./C.txt", quit, ch3, ch4, ch1, ch2)
	go wrt("./D.txt", quit, ch4, ch1, ch2, ch3)

	for i := 0; i < 4; i++ {
		<-quit
	}
}
