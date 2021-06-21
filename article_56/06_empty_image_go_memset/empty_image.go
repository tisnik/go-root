// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá šestá část
//    Programovací jazyk Go a assembler (3.část)
//    https://www.root.cz/clanky/programovaci-jazyk-go-a-assembler-3-cast/

package main

import (
	memset "github.com/tmthrgd/go-memset"
	"image"
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

func fillPixels(pixels []uint8) {
	memset.Memset(pixels, 0xff)
}

func main() {
	destinationImage := image.NewRGBA(image.Rect(0, 0, 256, 256))

	fillPixels(destinationImage.Pix)

	err := saveImage(DestinationImageFileName, destinationImage)
	if err != nil {
		log.Fatal(err)
	}
}
