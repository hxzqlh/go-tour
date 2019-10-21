package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ySlice := make([][]uint8, dy)
	for y := range ySlice {
		xSlice := make([]uint8, dx)
		for x := range xSlice {
			xSlice[x] = uint8(x % (y + 1))
		}
		ySlice[y] = xSlice
	}
	return ySlice
}

func main() {
	pic.Show(Pic)
}
