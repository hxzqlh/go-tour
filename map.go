package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	count := make(map[string]int)
	result := strings.Fields(s)
	for _, v := range result {
		count[v] += 1
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
