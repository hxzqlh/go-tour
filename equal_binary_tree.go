package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	"time"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		m := <-ch1
		n := <-ch2
		fmt.Printf("值：%d  %d\n", m, n)
		if m != n {
			return false
		}

	}
	return true

}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(2)))
	time.Sleep(50 * time.Millisecond)
	fmt.Println(Same(tree.New(2), tree.New(2)))
}
