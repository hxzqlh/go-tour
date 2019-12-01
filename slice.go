package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	row := make([][]uint8,dy)
	for i := range row{
		col := make([]uint8,dx)
		for j := range col{
			col[j] = uint8((i+j)/2)
		}
		row[i] = col
	}
	return row
}

func main() {
	pic.Show(Pic)
}
