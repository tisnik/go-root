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

func fillPixels(pixels []uint8)

func fillPixels2(pixels []uint8) {
	for i := 0; i < len(pixels); i++ {
		pixels[i] = 255
	}
}

func main() {
	destinationImage := image.NewRGBA(image.Rect(0, 0, 256, 256))

	println(destinationImage.Pix[0])
	println(destinationImage.Pix[1])
	println(destinationImage.Pix[2])
	fillPixels(destinationImage.Pix)
	println(destinationImage.Pix[0])
	println(destinationImage.Pix[1])
	println(destinationImage.Pix[2])

	err := saveImage(DestinationImageFileName, destinationImage)
	if err != nil {
		log.Fatal(err)
	}
}
