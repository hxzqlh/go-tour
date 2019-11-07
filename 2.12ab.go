package main

import (
	"fmt"
	"sync"
)

func main() {
	char := 'A'
	num := 1
	ch1 := make(chan bool, 1)
	ch1 <- true
	wg2 := sync.WaitGroup{}
	wg2.Add(2)
	go func() {
		defer wg2.Done()
		for i := 0; i < 1000; i++ {
			if t, ok := <-ch1; ok {
				if t {
					fmt.Print(num)
					num++
					fmt.Print(num)
					num++
				}
				ch1 <- false
			}
		}
	}()
	go func() {
		defer wg2.Done()
		for i := 0; i < 1000; i++ {
			if t, ok := <-ch1; ok {
				if !t {
					fmt.Printf("%c", char)
					char++
					fmt.Printf("%c", char)
					char++
					if char > 'Z' {
						char = 'A'
					}
				}
				ch1 <- true
			}
		}
	}()

	wg2.Wait()
}
