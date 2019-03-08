// Seriál "Programovací jazyk Go"
//
// Čtrnáctá část
//
// Demonstrační příklad číslo 1:
//     Vytvoření rastrového obrázku s výchozími hodnotami pixelů (zcela průhledné černé pixely)

package main

import (
	"image"
	"image/png"
	"os"
)

const width = 256
const height = 256

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	outfile, err := os.Create("01.png")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	png.Encode(outfile, img)
}
