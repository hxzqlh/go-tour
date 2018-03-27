package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Dfs(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Dfs(t.Left, ch)
	ch <- t.Value
	Dfs(t.Right, ch)
}

func Walk(t *tree.Tree, ch chan int) {
	Dfs(t, ch)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		switch {
		case ok1 != ok2:
			return false
		case ok1 == false:
			return true
		case v1 != v2:
			return false
		}
	}
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
