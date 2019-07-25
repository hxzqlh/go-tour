package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var pic [][]uint8
	for y := 0; y < dy; y++ {
		var l []uint8
		for x := 0; x < dx; x++ {
			l = append(l, (uint8)((x+y)/2))
		}
		pic = append(pic, l)
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
