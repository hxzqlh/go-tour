package main

import "golang.org/x/tour/pic"
func Pic(dx, dy int) [][]uint8 {
	x := 33
	y := 33
	board := make([][]uint8, dy)
	s := make([]uint8, dx)
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			s[j] = uint8((x^y))
			x+=(i*j)
			y+=(j/(i+1))
		}
		board[i] = s
	}
	return board
}

func main() {
	pic.Show(Pic)
}
