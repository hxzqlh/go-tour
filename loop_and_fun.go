package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x/2
	for math.Abs(z*z-x)>1e-10{
		z -= (z*z-x)/(2*z)
		fmt.Println("z:",z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
