package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(x/2);
	for i := 1; i <= 10; i++{
		z -= (z*z - x) / (2*z)
		fmt.Printf("第%v次 z=%v\n", i, z)
		//if z - math.Sqrt(x) == float64(0){
		//break
		//}
	}
	return z;
}

func main() {
	x := float64(3);
	pfg :=Sqrt(x);
	fmt.Printf("牛顿法实现的平方根为%v\n", pfg)
	fmt.Printf("标准库实现的平方根为%v\n", math.Sqrt(x))
	fmt.Printf("两者相差%v\n", pfg - math.Sqrt(x))
}
