package main

import (
	"fmt"
)

func main() {
	c1 := make(chan bool)
	c2 := make(chan bool)
	exit := make(chan int)

	go func() {
		for i := 1; i < 30; i += 2 {
			<-c2
			fmt.Print(i)
			fmt.Print(i + 1)
			c1 <- false
		}
	}()

	go func() {
		for i := 'A'; i < 'Z'; i += 2 {
			<-c1
			fmt.Printf("%c", i)
			fmt.Printf("%c", i+1)
			c2 <- true

		}
		exit <- 1
	}()

	c1 <- true
	<-exit
}
