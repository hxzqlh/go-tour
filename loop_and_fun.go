package main

import (
	"fmt"
	"time"
)

func charCreator(done chan bool) <-chan string {
	c := make(chan string)
	go func(done chan bool) {
		var startChar byte = 'A'
		s := string(startChar) + string(startChar+1)
		for {
			select {
			case c <- s:
				startChar = startChar + 2
				if startChar > 'Z' {
					startChar = 'A'
				}
				s = string(startChar) + string(startChar+1)
			case <-done:
				close(c)
				return
			}
		}
	}(done)

	return c
}

func numbCreator(done chan bool) <-chan string {
	c := make(chan string)
	go func(done chan bool) {
		var startInt = 1
		s := fmt.Sprintf("%v%v", startInt, startInt+1)
		for {
			select {
			case c <- s:
				startInt = startInt + 2
				if startInt > 10 {
					startInt = 1
				}
				s = fmt.Sprintf("%v%v", startInt, startInt+1)
			case <-done:
				close(c)
				return
			}
		}
	}(done)

	return c
}

func main() {
	done := make(chan bool)
	char := charCreator(done)
	numb := numbCreator(done)
	var exch = true

	timeout := time.After(time.Millisecond * 10)
	//指定时间超时之后，自动退出;
	go func(done chan bool) {
		select {
		case <-timeout:
			close(done)
		}
	}(done)

	for {
		if exch == true {
			n := <-numb
			if n == "" {
				fmt.Printf("\nChannel be closed!\n")
				break
			}
			fmt.Printf(n)
			exch = false
		} else {
			c := <-char
			if c == "" {
				fmt.Printf("\nChannel be closed!\n")
				break
			}
			fmt.Printf(c)
			exch = true
		}
	}

	fmt.Printf("\nGo out!\n")
	return
}

