package main

import (
	"fmt"
	"strconv"
	"time"
)

func printDig(ch chan int) {
	for i := 0; i < 100; i += 2 {
		ch <- i + 1
		ch <- i + 2
	}
}
func printChar(ch chan string) {
	for i := 0; i < 100; i += 2 {
		ch <- string((i % 26) + 'A')
		ch <- string(((i + 1) % 26) + 'A')
	}
}
func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	var result string
	var s string
	var n int
	go printDig(ch1)
	go printChar(ch2)
	time.Sleep(2 * time.Second)
	for i := 0; i < 200; i++ {
		switch i % 4 {
		case 0:
			n = <-ch1
			result += strconv.Itoa(n)
		case 1:
			n = <-ch1
			result += strconv.Itoa(n)
		case 2:
			s = <-ch2
			result += s
		case 3:
			s = <-ch2
			result += s
		}
	}
	fmt.Println(result)

}
