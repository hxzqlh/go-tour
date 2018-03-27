package main

import "fmt"

func fibonacci() func() int {
	var a, b, c = 0, 1, 1
	return func() int {
		ret := a
		a, b = b, c
		c = a + b
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
