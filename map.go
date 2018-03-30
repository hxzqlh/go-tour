package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	word := make(map[string]int)
	array := strings.Fields(s)
	for _, i := range array {
		word[i]++
	}
	return word
}

func main() {
	wc.Test(WordCount)
}
