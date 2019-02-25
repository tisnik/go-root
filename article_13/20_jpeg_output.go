package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

const width = 256
const height = 256

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			var red uint8 = uint8(x)
			var green uint8 = uint8((x - y))
			var blue uint8 = uint8(y)
			c := color.RGBA{red, green, blue, 255}
			img.SetRGBA(x, y, c)
		}
	}

	outfile, err := os.Create("test.jpeg")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	jpeg.Encode(outfile, img, nil)
}
