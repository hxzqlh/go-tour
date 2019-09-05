package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	var a0, a1 int
	a1 = 1
	return func() int {
		sum := a0
		a0, a1 = a1, (a0 + a1)
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
