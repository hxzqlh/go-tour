package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("cannot Sqrt negative number: %v",float64(e))
}

func Sqrt2(x float64) (float64, error) {
	if x > 0 {
		z := 1.0
		for i := 1; i <= 10; i++ {
			z -= (z*z - x) / (2 * z)
			if math.Abs((z*z-x)/(2*z)) < 0.00001 {
				return z,nil
			}
		}
	}
	return 0,ErrNegativeSqrt(x)
}

func main() {
	fmt.Println(Sqrt2(2))
	fmt.Println(Sqrt2(-2))
}
