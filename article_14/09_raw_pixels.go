package main

import (
	"image"
	"image/png"
	"os"
)

const width = 256
const height = 256

func main() {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	outfile, err := os.Create("09.png")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	for y := 0; y < height; y++ {
		index := img.Stride * y
		for x := 0; x < width; x++ {
			img.Pix[index] = 0
			index++
			img.Pix[index] = 0
			index++
			img.Pix[index] = 255
			index++
			img.Pix[index] = byte(x)
			index++
		}
	}
	png.Encode(outfile, img)
}
