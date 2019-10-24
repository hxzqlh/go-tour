package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		wordMap[word]++
	}
	return wordMap
}
func main() {
	wc.Test(WordCount)
}
