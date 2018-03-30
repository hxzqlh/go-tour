package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)
	ss := strings.Fields(s)
	for _, v := range ss {
		ret[v]++
	}
	return ret

}

func main() {
	wc.Test(WordCount)
}
