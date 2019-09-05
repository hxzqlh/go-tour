package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	count := make(map[string]int)
	words := strings.Fields(s)
	for _,v := range words {
		c, ok := count[v]
		if !ok {
			count[v] = 1
		}
		count[v] = c + 1
	}
	return count
}

func main() {
	wc.Test(WordCount)
}

