package main

import (
	"fmt"
	//"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	var temp float64 = x
	i := 1
	for {
		z -= (z*z - x) / (2 * z)
		if temp-z < 0.00000000001 {
			break
		}
		temp = z
		i++
	}
	fmt.Printf("µü´ú´ÎÊý£º%d\n", i)
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	//fmt.Println(math.Sqrt(2))
}
