package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	if e < 0 {
		return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
	}

	return ""
}

const EPS = 0.0000001

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	lastValue := 0.0
	i := 0
	for math.Abs(lastValue-z) > EPS {
		lastValue = z
		z -= (z*z - x) / (2 * z)
		i++
		fmt.Println(i, ":", z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

