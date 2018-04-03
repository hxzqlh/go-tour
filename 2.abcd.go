package main

import (
	"os"
	"strconv"
	"time"
)

var c [4]chan int

func print1234(i int, f [4]*os.File) {
	go func() {
		for n := 0; n < 100; n++ {
			<-c[i%4]
			for j := 0; j < 4; j++ {
				(*f[j]).WriteString(strconv.Itoa((i+j)%4+1) + " ")
			}
			c[(i+1)%4] <- i
		}
	}()
}

func creatFile(name string) *os.File {
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	return f
}

func main() {
	f := [4]*os.File{}
	f[0] = creatFile("A.txt")
	defer f[0].Close()
	f[1] = creatFile("B.txt")
	defer f[1].Close()
	f[2] = creatFile("C.txt")
	defer f[2].Close()
	f[3] = creatFile("D.txt")
	defer f[3].Close()

	for i := 0; i < 4; i++ {
		c[i] = make(chan int)
	}
	for j := 0; j < 4; j++ {
		print1234(j, f)
	}
	c[0] <- 1
	time.Sleep(500 * time.Millisecond)
}
