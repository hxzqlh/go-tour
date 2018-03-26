package main

import (
	"golang.org/x/tour/pic"
)

func Cal(x, y int) uint8 {
	return uint8((x + y) / 2)
}

func Pic(dx, dy int) [][]uint8 {
	matrix := make([][]uint8, dx)
	for i := range matrix {
		tmp := make([]uint8, dy)
		matrix[i] = tmp
		for j := range tmp {
			matrix[i][j] = Cal(i, j)
		}
	}
	return matrix
}

func main() {
	pic.Show(Pic)
}
