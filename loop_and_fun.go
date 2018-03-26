package main

import "fmt"

const EPS = 1e-10

func Fabs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func Sqrt(x float64) float64 {
	z := float64(1)
	cnt := 0
	for {
		dis := (z*z - x) / (2 * z)
		cnt++
		if Fabs(dis) < EPS {
			break
		}
		z -= dis
	}
	fmt.Println(cnt)
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
