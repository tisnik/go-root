// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá pátá část
//    Knihovna Gift pro zpracování rastrových obrázků v Go
//    https://www.root.cz/clanky/knihovna-gift-pro-zpracovani-rastrovych-obrazku-v-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z padesáté páté části:
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
const DestinationImageFileNameTemplate = "06_contrast_%d.png"

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

	for contrast := -50; contrast <= 50; contrast += 20 {
		g := gift.New(
			gift.Contrast(float32(contrast)))

		destinationImage := image.NewRGBA(g.Bounds(sourceImage.Bounds()))
		g.Draw(destinationImage, sourceImage)

		filename := fmt.Sprintf(DestinationImageFileNameTemplate, contrast)
		err = saveImage(filename, destinationImage)
		if err != nil {
			log.Fatal(err)
		}
	}
}
