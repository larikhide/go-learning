package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	Weight int
	Height int
}

func (im Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, im.Weight, im.Height)
}

func (im Image) ColorModel() color.Model {
	return color.ModelFunc(color.RGBAModel.Convert)
}

func (im Image) At(x, y int) color.Color {
	v := uint8(x*x + y*y + x*y/2*3*4)
	return color.RGBA{255, 255, v, v}
}

func main() {
	m := Image{255, 255}
	pic.ShowImage(m)
}
