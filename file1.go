package main

import (
	"fmt"
	// "io"
	// "io/ioutil"
	"os"
	"strconv"
)

func worker1(j chan int) {
	j <- 1
}
func worker2(j chan int) {
	j <- 2
}
func worker3(j chan int) {
	j <- 3
}
func worker4(j chan int) {
	j <- 4
}

func main() {
	a, err := os.Create("A")
	if err != nil {
		fmt.Errorf("Creat A failed")
	}
	b, err := os.Create("B")
	if err != nil {
		fmt.Errorf("Creat B failed")
	}
	c, err := os.Create("C")
	if err != nil {
		fmt.Errorf("Creat C failed")
	}
	d, err := os.Create("D")
	if err != nil {
		fmt.Errorf("Creat D failed")
	}
	j1 := make(chan int, 100)
	j2 := make(chan int, 100)
	j3 := make(chan int, 100)
	j4 := make(chan int, 100)
	go worker1(j1)
	go worker2(j2)
	go worker3(j3)
	go worker4(j4)

	w := <-j1
	x := <-j2
	y := <-j3
	z := <-j4
	str := strconv.Itoa(w) + " " + strconv.Itoa(x) + " " + strconv.Itoa(y) + " " + strconv.Itoa(z) + " "
	a.Write([]byte(str))
	str = strconv.Itoa(x) + " " + strconv.Itoa(y) + " " + strconv.Itoa(z) + " "
	b.Write([]byte(str))
	str = strconv.Itoa(y) + " " + strconv.Itoa(z) + " "
	c.Write([]byte(str))
	str = strconv.Itoa(z) + " "
	d.Write([]byte(str))
	num := 0
	for i := w; i < 5; i++ {
		str1 := strconv.Itoa(i) + " "
		a.Write([]byte(str1))
		str2 := strconv.Itoa(i) + " "
		b.Write([]byte(str2))
		str3 := strconv.Itoa(i) + " "
		c.Write([]byte(str3))
		str4 := strconv.Itoa(i) + " "
		d.Write([]byte(str4))
		if i == 4 {
			i = 0
		}
		num++
		if num > 10000 {
			break
		}
	}

}
