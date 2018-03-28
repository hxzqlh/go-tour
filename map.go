package main

import (
	"golang.org/x/tour/wc"
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	smap := make(map[string]int)
	str := strings.Fields(s)
	fmt.Println(str)
	for i := 0; i < len(str); i++ {
		smap[str[i]]++
	}
	return smap
}

func main() {
	wc.Test(WordCount)
}
