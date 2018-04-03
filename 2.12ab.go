package main

import (
	"fmt"
	"time"
)

var c1, c2 chan bool

func printNum() {
	for i := 0; ; i += 2 {
		<-c1
		fmt.Printf("%d%d", i%26+1, i%26+2)
		c2 <- true
	}
}

func printAlpha() {
	for i := 0; ; i += 2 {
		<-c2
		fmt.Printf("%c%c", i%26+65, i%26+66)
		c1 <- true
	}
}

func main() {
	c1 = make(chan bool)
	c2 = make(chan bool)

	go printNum()
	go printAlpha()
	c1 <- true
	time.Sleep(5 * time.Millisecond)
}
