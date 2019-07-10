package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	sli := strings.Fields(s)
	m := make(map[string]int, 0)
	for i := range sli{
		m[sli[i]]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
