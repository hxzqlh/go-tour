package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for i := 0; i < 10; i++ {
		x := <-c1
		y := <-c2
		if x != y {
			return false
		}
	}
	return true
}

func main() {
	c := make(chan int, 100)
	tr1 := tree.New(1)
	tr2 := tree.New(2)
	tr3 := tree.New(1)
	go Walk(tr1, c)
	for i := 0; i < 10; i++ {
		x := <-c
		fmt.Println(x)
	}
	fmt.Println(Same(tr1, tr2))
	fmt.Println(Same(tr1, tr3))
}
