package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	w int
	h int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(x, y int) color.Color {
	r := (uint8)((float64)(x) / (float64)(i.w) * 255.0)
	g := (uint8)((float64)(y) / (float64)(i.h) * 255.0)
	b := (uint8)((float64)(x*y) / (float64)(i.w*i.h) * 255.0)
	return color.RGBA{r, g, b, 255}
}

func main() {
	m := Image{255, 255}
	pic.ShowImage(m)
}
