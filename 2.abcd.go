package main

import (
	"fmt"
)

func main() {
	ch1, ch2, ch3, ch4 := make(chan int), make(chan int), make(chan int), make(chan int)
	go func() {
		for {
			ch1 <- 1
		}
	}()
	go func() {
		for {
			ch2 <- 2
		}
	}()
	go func() {
		for {
			ch3 <- 3
		}
	}()
	go func() {
		for {
			ch4 <- 4
		}
	}()

	slice1, slice2, slice3, slice4 := make([]int, 0), make([]int, 0), make([]int, 0), make([]int, 0)

	for i := 0; i < 4; i++ {
		slice1 = append(slice1, <-ch1, <-ch2, <-ch3, <-ch4)
		slice2 = append(slice2, <-ch2, <-ch3, <-ch4, <-ch1)
		slice3 = append(slice3, <-ch3, <-ch4, <-ch1, <-ch2)
		slice4 = append(slice4, <-ch4, <-ch1, <-ch2, <-ch3)
	}

	fmt.Print("A:")
	printSlice(slice1)
	fmt.Print("B:")
	printSlice(slice2)
	fmt.Print("C:")
	printSlice(slice3)
	fmt.Print("D:")
	printSlice(slice4)
}

func printSlice(numSlice []int) {
	for _, v := range numSlice {
		fmt.Printf(" %d", v)
	}
	fmt.Println("...")
}
