package main

import (
	"fmt"
)

//Ñ­»·10´Î
func Sqrt(x float64) float64 {
	z := float64(1)
	var n, m float64
	for i := 0; i <= 10; i++ {
		n = z*z - x
		m = n / (2 * z)
		z = z - m
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
