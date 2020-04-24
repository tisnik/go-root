package main

import (
	"image"
	"image/png"
	"log"
	"os"
)

const SourceImageFileName = "gopher.png"
const DestinationImageFileName = "gopher_copied_2.png"

func loadImage(filename string) (image.Image, error) {
	infile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer infile.Close()

	src, _, err := image.Decode(infile)
	if err != nil {
		return nil, err
	}
	return src, nil
}

func saveImage(filename string, img image.Image) error {
	outfile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outfile.Close()

	png.Encode(outfile, img)
	return nil
}

func copyPixels(src []uint8, dst []uint8) {
	copy(dst, src)
}

func main() {
	s, err := loadImage(SourceImageFileName)
	if err != nil {
		log.Fatal(err)
	}

	var sourceImage *image.RGBA = s.(*image.RGBA)
	destinationImage := image.NewRGBA(sourceImage.Bounds())

	copyPixels(sourceImage.Pix, destinationImage.Pix)

	err = saveImage(DestinationImageFileName, destinationImage)
	if err != nil {
		log.Fatal(err)
	}
}
