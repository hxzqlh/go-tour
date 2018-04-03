package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	result := map[string]int{}
	var word []string = strings.Split(s, " ")
	for i := 0; i < len(word); i++ {
		if _, ok := result[word[i]]; ok {
			result[word[i]]++
		} else {
			result[word[i]] = 1
		}
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
