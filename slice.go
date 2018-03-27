package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	result := [][]uint8{}
	for x := 0; x < dx; x++ {
		temp := make([]uint8, dy)
		for y := 0; y < dy; y++ {
			temp = append(temp, uint8(x*y))
		}
		result = append(result, temp)
	}
	return result
}

func main() {
	pic.Show(Pic)
}