package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 500, 500)
}

func (Image) At(x, y int) color.Color {
	v := (x * y) / 2
	return color.RGBA{uint8(v), uint8(v), 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
