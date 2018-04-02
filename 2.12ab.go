package main

import (
	"fmt"
	"time"
)

func print12(ch, other *chan int) {
	for i := 1; ; i += 2 {
		<- *ch
		fmt.Printf("%d%d", i, i+1)
		*other <- 1
	}
}

func printAB(ch, other *chan int) {
	for i := 'A'; ; i += 2 {
		if i > 'Z' {
			i = 'A'
		}
		<- *ch
		fmt.Printf("%c%c", i, (byte)(i+1))
		*other <- 1
	}
}

func main() {
	a := make(chan int)
	b := make(chan int)
	go print12(&a, &b)
	go printAB(&b, &a)
	a <- 1
	time.Sleep(100*time.Millisecond)
}