package main

import "fmt"
import "math"

func Sqrt1(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println("z = ", z)
	}
	return z
}
func main() {
	fmt.Println(Sqrt1(2))
	fmt.Println(math.Sqrt(2))
}
