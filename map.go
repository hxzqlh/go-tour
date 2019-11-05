package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	res := make(map[string]int)
	arr := strings.Fields(s)
	for _, word := range arr {
		res[word] += 1
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
