package main

import (
	"fmt"
)

func worker1(num chan int) {
	for i := 1; i < 200; i++ {
		num <- i
	}
	close(num)
	return
}

func worker2(str chan string) {
	str1 := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < 26; i++ {
		str <- string(str1[i])
		if i == 25 {
			i = -1
		}
	}
}
func main() {
	num := make(chan int, 100)
	str := make(chan string, 100)

	go worker1(num)
	go worker2(str)

	for i := 0; i < 50; i++ {
		n1 := <-num
		n2 := <-num
		str1 := <-str
		str2 := <-str
		fmt.Printf("%d%d%s%s", n1, n2, str1, str2)
	}

	return
}
