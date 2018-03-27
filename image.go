package main

import (
	"golang.org/x/tour/pic"
	"image/color"
	"image"
)

type Image struct{
	w,h int
}

func (im Image)ColorModel() color.Model {
	return color.RGBAModel
}

func (im Image)Bounds() image.Rectangle {
	return image.Rect(0,0, im.w, im.h)
}

func (im Image) At(x, y int) color.Color  {
	return color.RGBA{uint8(x),uint8(y), 255, 255}
}

func main() {
	m := Image{w:10,h:10}
	pic.ShowImage(m)
}
