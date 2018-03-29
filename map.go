package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var m = make(map[string]int)
	for _, stemp := range strings.Fields(s) {
		a, _ := m[stemp]
		m[stemp] = a + 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
