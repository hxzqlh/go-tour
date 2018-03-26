package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	mp := map[string]int{}
	for _, str := range words {
		if _, in := mp[str]; in {
			mp[str]++
		} else {
			mp[str] = 1
		}
	}
	return mp
}

func main() {
	wc.Test(WordCount)
}
