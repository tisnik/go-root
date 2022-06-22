// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá pátá část
//    Knihovna Gift pro zpracování rastrových obrázků v Go
//    https://www.root.cz/clanky/knihovna-gift-pro-zpracovani-rastrovych-obrazku-v-go/
//
// Seznam příkladů z padesáté páté části:
//    https://github.com/tisnik/go-root/blob/master/article_55/README.md

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
const DestinationImageFileNameTemplate = "08_gamma_%4.2f.png"

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

	for gamma := 0.25; gamma <= 4.0; gamma *= 2 {
		g := gift.New(
			gift.Grayscale(),
			gift.Gamma(float32(gamma)))

		destinationImage := image.NewRGBA(g.Bounds(sourceImage.Bounds()))
		g.Draw(destinationImage, sourceImage)

		filename := fmt.Sprintf(DestinationImageFileNameTemplate, gamma)
		err = saveImage(filename, destinationImage)
		if err != nil {
			log.Fatal(err)
		}
	}
}
