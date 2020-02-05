package main

import (
	"github.com/disintegration/gift"
	"image"
	"image/png"
	"log"
	"os"
)

const SourceImageFileName = "Lenna.png"
const DestinationImageFileName = "15_sharpen.png"

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

func main() {
	sourceImage, err := loadImage(SourceImageFileName)
	if err != nil {
		log.Fatal(err)
	}

	filter := gift.Convolution([]float32{
		0, -1, 0,
		-1, 5, 1,
		0, -1, 0,
	}, true, false, false, 0.0)

	g := gift.New(filter)

	destinationImage := image.NewRGBA(g.Bounds(sourceImage.Bounds()))
	g.Draw(destinationImage, sourceImage)

	err = saveImage(DestinationImageFileName, destinationImage)
	if err != nil {
		log.Fatal(err)
	}
}
