package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	x,y,n:=0,1,0
	return func() int {
		n = x
		x,y = y ,x+y
		return n
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
