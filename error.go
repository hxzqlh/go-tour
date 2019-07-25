package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	errorinfo := fmt.Sprintf("cannot Sqrt negative number: %x", float64(e))
	return errorinfo
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		fmt.Println(ErrNegativeSqrt(x).Error())

	} else {
		z := 1.0
		for i := 0; i < 10; i++ {
			z -= (z*z - x) / (2 * z)
			fmt.Printf("%x\n", z)
		}
		return z, nil
	}
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
