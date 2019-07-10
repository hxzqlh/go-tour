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
func Same(t1, t2 *tree.Tree)bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go Walk(t1,ch1)
	go Walk(t2,ch2)
	for i:=0;i<len(ch1) ;i++ {
		if <-ch1 != <- ch2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int, 10)
	go Walk(tree.New(2), ch)
	for i:=0;i<10 ;i++ {
		fmt.Printf("%d ",<-ch)
	}
	fmt.Println()

	fmt.Println("两树储存的值相等？",Same(tree.New(1), tree.New(1)))
}
