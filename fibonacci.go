package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	a, b, c := 0, 1, 2
	return func(x int) int {
		if x < 2 {
			return x
		} else {
			c = a + b
			a = b
			b = c
			return c
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
