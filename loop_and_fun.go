package main

import (
	"fmt"
)

func Sqrt1(x float64) float64 {
	var z float64 = 1
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	Sqrt1(2)
}
