package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.5
	i := 0
	var d float64
	for {
		d = z
		z -= (z*z - x) / (2 * z)
		i++
		if (z-d) < 0.000000001 && (z-d) > 0 || (z-d) > -0.000000001 && (z-d) < 0 {
			fmt.Printf("执行了%d次\n", i)
			return z
		}

	}
}

func main() {
	fmt.Println(Sqrt(3))
}
