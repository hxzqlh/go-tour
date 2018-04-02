package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	s1 := make([][]uint8, dy)
	for x := range s1 {
		s2 := make([]uint8, dx)
		for y := range s2 {
			s2[y] = uint8((x + y) / 2)
		}
		s1[x] = s2
	}
	return s1
}

func main() {
	pic.Show(Pic)
}
