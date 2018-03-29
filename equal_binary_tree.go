package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int, flag bool) {

	if t.Left != nil {
		Walk(t.Left, ch, false)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch, false)
	}
	if flag {
		close(ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1, true)
	go Walk(t2, ch2, true)
	for v1 := range ch1 {
		if v1 != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	/*ch := make(chan int)
	go Walk(tree.New(1), ch, true)
	for v := range ch{
		fmt.Println(v)
	}*/
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
