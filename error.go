package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		var err ErrNegativeSqrt
		err = ErrNegativeSqrt(x)
		return 0, err
	} else {
		var z float64 = x / 2
		var n float64 = (z*z - x) / (2 * z)
		for n*n > (0.000000000000001) {
			n = (z*z - x) / (2 * z)
			z -= n
		}
		return z, nil
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
