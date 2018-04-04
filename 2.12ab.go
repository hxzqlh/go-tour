package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 26; i += 2 {
		fmt.Printf("%d%d", i+1, i+2)
		fmt.Printf("%c%c", 'A'+i, 'A'+i+1)
	}
}
