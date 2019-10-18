package main

import "golang.org/x/tour/pic"

//练习的题目描述的不清楚，不知道在讲什么
func Pic(dx, dy int) [][]uint8 {
	var slc [][]uint8
	for i := 0; i < dy; i++ {
		slc = append(slc, make([]uint8, dx))
	}
	return slc
}

func main() {
	pic.Show(Pic)
}
