package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wc := make(map[string]int)
	words := strings.Fields(s)
	for _, w := range words {
		wc[w]++
	}
	return wc
}

func main() {
	wc.Test(WordCount)
}
