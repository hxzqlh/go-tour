package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	sum, sum1, sum2 := 0, 0, -1
	//	fmt.Println(sum, sum1, sum2)
	return func() int {
		if sum2 == -1 {
			sum2++
			return sum2
		} else if sum2 == 0 {
			sum2++
			return sum2
		} else {
			sum = sum1 + sum2
			sum1 = sum2
			sum2 = sum
			return sum2
		}
		//		return 6
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
