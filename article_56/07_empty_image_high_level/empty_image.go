package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

const DestinationImageFileName = "empty.png"

func saveImage(filename string, img image.Image) error {
	outfile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outfile.Close()

	png.Encode(outfile, img)
	return nil
}

func fillPixels(img *image.RGBA) {
	clr := color.RGBA{255, 255, 255, 255}
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.SetRGBA(x, y, clr)
		}
	}
}

func main() {
	destinationImage := image.NewRGBA(image.Rect(0, 0, 256, 256))

	fillPixels(destinationImage)

	err := saveImage(DestinationImageFileName, destinationImage)
	if err != nil {
		log.Fatal(err)
	}
}
