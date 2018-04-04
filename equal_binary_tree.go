package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk 遍历 tree t 将所有的值从 tree 发送到 channel ch。
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
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	t := tree.New(1)
	go Walk(t, ch)
	for i := 0; i < 10; i++ {
		fmt.Println("Got value:", <-ch)
	}

	fmt.Println("Check same:", Same(tree.New(1), tree.New(1)))
	fmt.Println("Check same:", Same(tree.New(1), tree.New(2)))
}
