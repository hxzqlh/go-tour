// 2.abcd project main.go
package main

import (
	"fmt"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(4)
	file1, _ := os.Create("C:/Users/11326/Desktop/go-tour/go-tour/A")
	file2, _ := os.Create("C:/Users/11326/Desktop/go-tour/go-tour/B")
	file3, _ := os.Create("C:/Users/11326/Desktop/go-tour/go-tour/C")
	file4, _ := os.Create("C:/Users/11326/Desktop/go-tour/go-tour/D")
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)
	go print(ch1, f1)
	go print(ch2, f2)
	go print(ch3, f3)
	go print(ch4, f4)
	ch1 <- 0
	ch2 <- 1
	ch3 <- 2
	ch4 <- 3
	wg.Wait()
	file1.Close()
	file2.Close()
	file3.Close()
	file4.Close()
}
func print(ch chan int, f *os.File) {
	num := <-ch
	t := (num % 4) + 1
	f.Write([]byte(fmt.Sprint(t)))
	num += 1
	if num == 20 {
		wg.Done()
		close(ch)
		return
	}
	go print(ch, f)
	ch <- num
}
