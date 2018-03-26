package main

import "fmt"

const EPS2 = 1e-10

func Fabs2(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func Sqrt2(x float64) (float64, error) {
	z := float64(1)
	var err error
	if x < 0 {
		err = ErrNegativeSqrt(-2)
		return 0, err
	}
	cnt := 0
	for {
		dis := (z*z - x) / (2 * z)
		cnt++
		if Fabs2(dis) < EPS2 {
			break
		}
		z -= dis
	}
	fmt.Println(cnt)
	return z, err
}

func main() {
	fmt.Println(Sqrt2(2))
	fmt.Println(Sqrt2(-2))
}
