package main

import "fmt"

func main() {
	c1, c2 := make(chan int), make(chan int)
	go func() {
		a, b := 1, 65
		for {
			select {
			case <-c1:
				fmt.Print(a)
				a++
			case <-c2:
				fmt.Print(string(b))
				b++
			}
		}
	}()
	for i := 1; i < 27; i++ {
		if i%2 != 0 {
			c1 <- 1
			c1 <- 1
		} else {
			c2 <- 2
			c2 <- 2
		}
	}
}
