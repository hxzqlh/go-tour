package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("this error num is %v", float64(e))
}

func Sqrt1(x float64) (float64, error) {
	const Accuracy = 0.0000001
	v := 5.0
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	for i := 0.0; math.Abs(v-i) > Accuracy; {
		i = v
		v -= (v*v - x) / (2 * x)
		fmt.Println(v)
	}
	return v, nil
}

func main() {
	fmt.Println(Sqrt1(-10))
	fmt.Println(Sqrt1(10))
}
