package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	a := make(map[string]int)
	var j int
	b := strings.Fields(s)
	for i, value := range b {
		a[value]++
		j = i
	}
	fmt.Printf("一共有%d个单词\n", j+1)
	return a
}

func main() {
	wc.Test(WordCount)

}
