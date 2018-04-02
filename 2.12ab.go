/*2.交替打印数字和字母，使得最终效果如下：
12AB34CD56EF78GH910IJ...*/
package main

import (
	"fmt"
)

var ch1 chan int
var ch2 chan byte
var b int
var c byte

func main() {
	ch1 = make(chan int)
	ch2 = make(chan byte)
	go Ch1()
	go Ch2()
	for i := 0; i < 20; i++ {
		num := i % 4
		if num < 2 {
			fmt.Print(<-ch1)
		} else {
			fmt.Print(string(<-ch2))
		}
	}
}

func Ch1() {
	b = 1
	for {
		ch1 <- b
		b++
	}
}

func Ch2() {
	c = 65
	for {
		ch2 <- c
		c++
	}
}
