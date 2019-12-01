package main

import(
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{
	width int
	height int
}

func (img Image) Bounds() image.Rectangle{
	return image.Rect(0,0,img.width,img.height)
}
func (img Image) ColorModel() color.Model{
	return color.RGBAModel
}
func (img Image) At(x,y int) color.Color{
	return color.RGBA{uint8(x),uint8(y),255,255}
}

func main() {
	m := Image{100,100}
	pic.ShowImage(m)
}
