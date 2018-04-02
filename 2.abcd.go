package main

import (
	"fmt"
	"time"
	"os"
	"bufio"
)

type ch chan int

func print1234(n int, ch, nxt *ch, w *bufio.Writer) {
	for {
		<- *ch
		w.WriteString(fmt.Sprintf("%d ", n))
		*nxt <- 1
	}
}

func Print1234(l, start int, file string, mp *map[int]int) {
	f, err := os.Create(file)
	defer f.Close()
	if err != nil {
		fmt.Println("Error, can not create the file")
		return
	}
	w := bufio.NewWriter(f)
	chs := make([]ch, l)
	for i := range chs {
		chs[i] = make(chan int)
	}
	for i := 0; i < l; i++ {
		go print1234(i+1, &chs[i], &chs[(*mp)[i+1]-1], w)
	}
	chs[start-1] <- 1
	time.Sleep(100*time.Millisecond)
	w.Flush()
}

func main() {
	mp := map[int]int{
		1:2,
		2:3,
		3:4,
		4:1,
	}
	Print1234(4, 1, "A", &mp)
	Print1234(4, 2, "B", &mp)
	Print1234(4, 3, "C", &mp)
	Print1234(4, 4, "D", &mp)
}