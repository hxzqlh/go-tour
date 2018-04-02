package main

import (
	"fmt"
)

func main() {
	chNum := make(chan int, 2)
	chAlpha := make(chan string, 2)
	countNum := 0
	countAlpha := 0

	num := 1
	alpha := 65

	go func() {
		for {
			if countNum == 13 {
				num = 1
			}
			chNum <- num
			num++
			chNum <- num
			num++
			countNum++
		}
	}()

	go func() {
		for {
			if countAlpha == 13 {
				alpha = 65
			}
			chAlpha <- string(alpha)
			alpha = alpha + 1
			chAlpha <- string(alpha)
			alpha = alpha + 1
			countAlpha++
		}
	}()

	for {
		fmt.Print(<-chNum)
		fmt.Print(<-chNum)
		fmt.Print(<-chAlpha)
		fmt.Print(<-chAlpha)
	}

}
