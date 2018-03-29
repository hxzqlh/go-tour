package main

import (
	"golang.org/x/tour/pic"
	//"math"
)

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)
	for y, _ := range a {
		a[y] = make([]uint8, dx)
		for x, _ := range a[y] {
			a[y][x] = uint8(x % (y + 1))
		}
	}
	return a
}

func main() {
	pic.Show(Pic)
}
