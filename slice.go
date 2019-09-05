package main

import (
	"golang.org/x/tour/pic"
	//"math"
)

func Pic(dx, dy int) [][]uint8 {
	var pic [][]uint8

	for i := 0; i < dy; i++ {
		var tmp []uint8
		for j := 0; j < dx; j++ {
			tmp = append(tmp, uint8(j%(i+1)))
		}
		pic = append(pic, tmp)
	}

	return pic
}

func main() {
	pic.Show(Pic)
}

