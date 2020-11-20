// Seriál "Programovací jazyk Go"
//
// Padesátá pátá část
//    Knihovna Gift pro zpracování rastrových obrázků v Go
//    https://www.root.cz/clanky/knihovna-gift-pro-zpracovani-rastrovych-obrazku-v-go/

package main

import (
	"github.com/disintegration/gift"
	"image"
	"image/png"
	"log"
	"os"
)

const SourceImageFileName = "Lenna.png"
const DestinationImageFileName = "03_invert.png"

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

	g := gift.New(
		gift.Invert())

	destinationImage := image.NewRGBA(g.Bounds(sourceImage.Bounds()))
	g.Draw(destinationImage, sourceImage)

	err = saveImage(DestinationImageFileName, destinationImage)
	if err != nil {
		log.Fatal(err)
	}
}
