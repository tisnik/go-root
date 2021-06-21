// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá šestá část
//    Programovací jazyk Go a assembler (3.část)
//    https://www.root.cz/clanky/programovaci-jazyk-go-a-assembler-3-cast/

package main

import (
	"image"
	"image/png"
	"log"
	"os"
)

const SourceImageFileName = "gopher.png"
const DestinationImageFileName = "gopher_copied.png"

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
	for i := 0; i < len(src); i++ {
		dst[i] = src[i]
	}
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
