package main

import "fmt"

func Fibonacci() func() int {
	num1 := 0
	num2 := 1
	return func() int {
		temp := num1
		num1 = num2
		num2 = num1 + num2
		return temp
	}
}

func main() {
	f := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
