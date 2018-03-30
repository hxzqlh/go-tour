package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	count := make(map[string]int)

	for _, value := range strings.Fields(s) {
		count[value]++
	}

	return count
}

func main() {
	wc.Test(WordCount)
}
