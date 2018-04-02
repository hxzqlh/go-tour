package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	y := 0.0
	for {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
		if math.Abs(z-y) < 0.00000000000001 {
			break
		} else {
			y = z
		}
	}
	return z
}

func main() {
	fmt.Println("牛顿算法的值", Sqrt(2))
	fmt.Println("mathsqrt的值", math.Sqrt(2))
}
