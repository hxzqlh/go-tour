package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	// first := 1
	// second := 0
	// temp := 0
	// fmt.Println(0)
	// return func() int{
	// 	temp = second
	// 	second += first
	// 	first = temp
	// 	return second
	// }
	fmt.Println(0)
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return y
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
