package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		picture[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			picture[i][j] = uint8(dx * dy)
		}
	}
	return picture
}

func main() {
	pic.Show(Pic)
}
