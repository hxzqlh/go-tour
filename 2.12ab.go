package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	numChan := make(chan int)
	strChan := make(chan string)
	quit := make(chan bool)

	go func() {
		for i := 1; i <= 27; i += 2 {
			<-strChan
			fmt.Printf("%d", i)
			fmt.Printf("%d", i+1)
			numChan <- 1
		}
	}()
	go func() {
		for i := 'A'; i < 'Z'; i += 2 {
			<-numChan
			fmt.Printf("%c", i)
			fmt.Printf("%c", i+1)
			strChan <- "go"
		}
		quit <- true
	}()

	strChan <- "go"
	<-quit
}
