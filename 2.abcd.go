package main

import (
	"fmt"
	"time"
)

func groutine(p chan int) (int, int, int, int) {
	go chanA(p)
	a := <-p
	go chanB(p)
	b := <-p
	go chanC(p)
	c := <-p
	go chanD(p)
	d := <-p
	return a, b, c, d
}
func chanA(p chan int) {
	a := 1
	p <- a
}

func chanB(p chan int) {
	b := 2
	p <- b
}

func chanC(p chan int) {
	c := 3
	p <- c
}

func chanD(p chan int) {
	d := 4
	p <- d
}
func main() {
	msg := make(chan int)
	a, b, c, d := groutine(msg)
	fmt.Printf("A:")
	for i := 1; i < 5; i++ {
		fmt.Printf(" %d %d %d %d", a, b, c, d)
	}
	fmt.Printf("\n")
	fmt.Printf("B:")
	for i := 1; i < 5; i++ {
		fmt.Printf(" %d %d %d %d", b, c, d, a)
	}
	fmt.Printf("\n")
	fmt.Printf("C:")
	for i := 1; i < 5; i++ {
		fmt.Printf(" %d %d %d %d", c, d, a, b)
	}
	fmt.Printf("\n")
	fmt.Printf("D:")
	for i := 1; i < 5; i++ {
		fmt.Printf(" %d %d %d %d", d, a, b, c)
	}
	time.Sleep(time.Second * 1)

}
