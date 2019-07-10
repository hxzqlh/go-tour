package main

import (
	"fmt"
)

func main() {
	ch1, ch2, ch3, ch4 := make(chan int), make(chan int), make(chan int), make(chan int)
	go func() { for{ch1 <- 1} }()
	go func() { for{ch2 <- 2 }}()
	go func() { for{ch3 <- 3 }}()
	go func() { for{ch4 <- 4 }}()

	sli1, sli2, sli3, sli4 := make([]int, 0), make([]int, 0), make([]int, 0), make([]int, 0)
	func() {for i:=0;i<4;i++{ sli1 = append(sli1, <-ch1, <-ch2, <-ch3, <-ch4)} }()
	func() {for i:=0;i<4;i++{ sli2 = append(sli2, <-ch2, <-ch3, <-ch4, <-ch1)} }()
	func() {for i:=0;i<4;i++{ sli3 = append(sli3, <-ch3, <-ch4, <-ch1, <-ch2)} }()
	func() {for i:=0;i<4;i++{ sli4 = append(sli4, <-ch4, <-ch1, <-ch2, <-ch3)} }()


	fmt.Print("A:")
	print(sli1)
	fmt.Print("B:")
	print(sli2)
	fmt.Print("C:")
	print(sli3)
	fmt.Print("D:")
	print(sli4)
}
//遍历切片并打印
func print(sli []int)  {
	for _,v:=range sli{
		fmt.Printf(" %d",v)
	}
	fmt.Println(" …")
}