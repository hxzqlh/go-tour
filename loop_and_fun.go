package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	const Accuracy = 0.0000001
	v := 5.0
	for i := 0.0; math.Abs(v-i) > Accuracy; {
		i = v
		v -= (v*v - x) / (2 * x)
		fmt.Println(v)
	}
	return v
}

func main() {
	fmt.Println(Sqrt(10))
	fmt.Println(math.Sqrt(10))
}
