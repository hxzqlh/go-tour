package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int, depth int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch, depth+1)
	ch <- t.Value
	Walk(t.Right, ch, depth+1)

	if depth == 0 {
		close(ch)
	}

	return
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1, 0)
	go Walk(t2, ch2, 0)

	for i := 0; i < 10; i++ {
		v1 := <-ch1
		v2 := <-ch2
		if v1 != v2 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(Same(tree.New(3), tree.New(3)))
}

