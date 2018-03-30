package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	test := make(map[string]int)
	num := len(strings.Fields(s))
	temp := make([]string, num)
	temp = change(strings.Fields(s))
	d := 1
leb1:
	for i := 0; i < num; i++ {
		if i+1 < num && temp[i] == temp[i+1] {
			d++
			continue leb1
		}
		test[temp[i]] = d
		d = 1
	}
	return test

}
func change(t []string) []string {
	for i := 0; i < len(t); i++ {
		for j := i + 1; j < len(t); j++ {
			if t[i] > t[j] {
				a := t[i]
				t[i] = t[j]
				t[j] = a
			}
		}
	}
	return t
}

func main() {
	wc.Test(WordCount)
}
