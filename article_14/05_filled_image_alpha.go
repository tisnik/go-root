package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

const width = 256
const height = 256

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	outfile, err := os.Create("05.png")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	for x := 0; x < width; x++ {
		alpha := byte(x)
		c := color.RGBA{0, 255, 0, alpha}
		for y := 0; y < height; y++ {
			img.SetRGBA(x, y, c)
		}
	}
	png.Encode(outfile, img)
}
