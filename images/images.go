package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	Rect image.Rectangle
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) Bounds() image.Rectangle {
	return i.Rect
}

func (i *Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(x / 2), 255, 255}
}

func main() {
	m := &Image{image.Rect(0, 0, 200, 300)}
	pic.ShowImage(m)
}
