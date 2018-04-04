package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	x, y int
}

func (p *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, (*p).x, (*p).y)
}

func (p *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (p *Image) At(x, y int) color.Color {

	return color.RGBA{uint8(x), uint8(y), uint8(255), uint8(255)}
}

func main() {

	m := &Image{150, 150}
	pic.ShowImage(m)
}
