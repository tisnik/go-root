package main

import (
	"image"
)

const width = 256
const height = 256

func main() {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	println("Stride: ", img.Stride)
	println("[]byte: ", len(img.Pix))
	r := img.Rect
	println("Rectangle:")
	println("    point 1: ", r.Min.X, r.Min.Y)
	println("    point 2: ", r.Max.X, r.Max.Y)
}
