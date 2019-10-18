package main

import "fmt"

// 返回一个“返回int的函数” 怎么用defer
func fibonacci() func() int {
	x, y := -1, 1
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
