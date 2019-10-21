package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	fmt.Println("Input num is", x)
	z := 1.0
	for i := 1; i <= 10; i++ {
		temp := z - (z*z-x)/(2*z)
		if temp == z || math.Abs(z-temp) < 0.00000000001 {
			break
		}
		z = temp
		fmt.Println(i, z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(1))
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(3))
}
