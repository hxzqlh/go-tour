package main

import (
	"fmt"
	"math"
)

const EPS = 0.0000001

func Sqrt(x, initial float64) float64 {
	z := initial
	lastValue := 0.0
	i := 0
	for math.Abs(lastValue-z) > EPS {
		lastValue = z
		z -= (z*z - x) / (2 * z)
		i++
		fmt.Println(i, ":", z)
	}
	return z
}

func main() {
	i := 23.0
	fmt.Println("math.Sqrt of", i, "is", math.Sqrt(float64(i)))

	fmt.Println("Sqrt of", i, "is", Sqrt(float64(i), 1.0))

	fmt.Println("Sqrt of", i, "is", Sqrt(float64(i), 2.0))

	fmt.Println("Sqrt of", i, "is", Sqrt(float64(i), 3.0))
}
