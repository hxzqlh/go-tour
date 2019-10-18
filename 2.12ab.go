package main

import "fmt"

func prtNum(ch chan [2]int) {
	for {
		for i := 0; i < 26; i += 2 {
			ch <- [2]int{i + 1, i + 2}
		}
	}
	// close(ch)
}

func prtLet(ch chan [2]rune) {
	for {
		for i := 0; i < 26; i += 2 {
			ch <- [2]rune{rune(i + 65), rune(i + 66)}
		}
	}
	// close(ch)
}

func main() {
	chnum := make(chan [2]int)
	chlet := make(chan [2]rune)
	go prtNum(chnum)
	go prtLet(chlet)
	for i := 0; i < 26; i += 2 {
		n, l := <-chnum, <-chlet
		fmt.Printf("%d%d%s%s", n[0], n[1], string(l[0]), string(l[1]))
	}
}
