package main

import (
    "image"
    "image/color"

    "golang.org/x/tour/pic"
)

type Image struct {
    W int
    H int
}

func (i Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, i.W, i.H)
}

func (i Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (self Image) At(x, y int) color.Color {
    return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
    m := Image{200, 200}
    pic.ShowImage(m)
}
