package main

import (
	"fmt"
	"github.com/disintegration/gift"
	"image"
	"image/png"
	"log"
	"os"
)

const SourceImageFileName = "Lenna.png"
const DestinationImageFileNameTemplateMin = "11_min_%02d.png"
const DestinationImageFileNameTemplateMax = "11_max_%02d.png"
const DestinationImageFileNameTemplateMean = "11_mean_%02d.png"
const DestinationImageFileNameTemplateMedian = "11_median_%02d.png"

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

	for size := 2; size <= 16; size *= 2 {
		g := gift.New(gift.Minimum(size, false))
		destinationImage := image.NewRGBA(g.Bounds(sourceImage.Bounds()))
		g.Draw(destinationImage, sourceImage)
		filename := fmt.Sprintf(DestinationImageFileNameTemplateMin, size)
		err = saveImage(filename, destinationImage)
		if err != nil {
			log.Fatal(err)
		}

		g = gift.New(gift.Maximum(size, false))
		destinationImage = image.NewRGBA(g.Bounds(sourceImage.Bounds()))
		g.Draw(destinationImage, sourceImage)
		filename = fmt.Sprintf(DestinationImageFileNameTemplateMax, size)
		err = saveImage(filename, destinationImage)
		if err != nil {
			log.Fatal(err)
		}

		g = gift.New(gift.Mean(size, false))
		destinationImage = image.NewRGBA(g.Bounds(sourceImage.Bounds()))
		g.Draw(destinationImage, sourceImage)
		filename = fmt.Sprintf(DestinationImageFileNameTemplateMean, size)
		err = saveImage(filename, destinationImage)
		if err != nil {
			log.Fatal(err)
		}

		g = gift.New(gift.Median(size, false))
		destinationImage = image.NewRGBA(g.Bounds(sourceImage.Bounds()))
		g.Draw(destinationImage, sourceImage)
		filename = fmt.Sprintf(DestinationImageFileNameTemplateMedian, size)
		err = saveImage(filename, destinationImage)
		if err != nil {
			log.Fatal(err)
		}

	}
}
