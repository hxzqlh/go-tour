package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	var z float64 = x / 2
	var n float64 = (z*z - x) / (2 * z)
	for n*n > (0.000000000000001) {
		n = (z*z - x) / (2 * z)
		z -= n
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
