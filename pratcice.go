package main

import (
	"fmt"
	"time"
)

func first() {
	a := 0
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Millisecond)
		if a == 4 {
			a = 1
		} else {
			a++
		}
		fmt.Print(a)
	}
	fmt.Println("")
}
func second() {
	b := 1
	for i := 0; i < 10; i++ {
		time.Sleep(15 * time.Millisecond)
		if b == 4 {
			b = 1
		} else {
			b++
		}
		fmt.Print(b)

	}
	fmt.Println("")
}
func third() {
	c := 2
	for i := 0; i < 10; i++ {
		time.Sleep(200 * time.Millisecond)
		if c == 4 {
			c = 1
		} else {
			c++
		}
		fmt.Print(c)

	}
	fmt.Println("")
}
func forth() {
	d := 3
	for i := 0; i < 10; i++ {
		time.Sleep(2500 * time.Millisecond)
		if d == 4 {
			d = 1
		} else {
			d++
		}
		fmt.Print(d)

	}
	fmt.Println("")
}
func main() {
	go first()
	go second()
	go third()
	go forth()
	time.Sleep(30 * time.Second)
}
