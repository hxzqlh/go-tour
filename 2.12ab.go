package main

import (
	"fmt"
	"sync"
)

func PrintNums(printChar chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 0; i < 5; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d", 2*i+j+1)
		}

		printChar <- 1
		<-printChar
	}
}

func PrintChars(printChar chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 0; i < 5; i++ {

		<-printChar

		for j := 0; j < 2; j++ {
			fmt.Printf("%c", 'A'+(2*i+j))
		}

		printChar <- 1
	}
}

func main() {
	var wg sync.WaitGroup
	flag := make(chan int)
	wg.Add(2)

	go PrintNums(flag, &wg)
	go PrintChars(flag, &wg)

	wg.Wait()
}
