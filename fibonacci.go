package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	cur := 0
	next := 1
	fun := func() int {
		tmp := cur
		cur = next
		next += tmp
		return tmp
	}
	return fun
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
