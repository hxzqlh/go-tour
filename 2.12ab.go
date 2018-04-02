package main

import (
	"fmt"
	"sync"
)

func PrintNums(printChar chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 0; i < 50; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d", 2*i+j+1)
		}

		printChar <- 1
		<-printChar
	}
}

func PrintChars(printChar chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 0; i < 50; i++ {
		temp := i % 13
		<-printChar

		for j := 0; j < 2; j++ {
			fmt.Printf("%c", 'A'+(2*temp+j))
		}

		printChar <- 1
	}
}

func main() {

	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go PrintNums(ch, &wg)
	go PrintChars(ch, &wg)

	wg.Wait()
}
