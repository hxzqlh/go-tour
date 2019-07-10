package main

import (
	"fmt"
	"math"
)

func Sqrt1(x float64) float64 {
	z := 1.0
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(i,z)
		if math.Abs((z*z-x)/(2*z)) < 0.00001 {
			return z
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt1(5))
}
