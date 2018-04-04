package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	ss := strings.Fields(s)
	m := make(map[string]int)
	for i := 0;i < len(ss);i++{
		m[ss[i]] = strings.Count(s,ss[i])//用strings的Count函数来计数
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
