package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		c := a
		a = b
		b = c + b
		return c
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 5; i++ {
		fmt.Println(f())
	}
}
