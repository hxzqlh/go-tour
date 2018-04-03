package main

import "fmt"

func main() {
	for i := 0; i < 1000; i += 2 {
		fmt.Printf("%d%d%s%s", i+1, i+2, string(byte(i%26)+'A'), string(byte((i+1)%26)+'A'))
	}
}
