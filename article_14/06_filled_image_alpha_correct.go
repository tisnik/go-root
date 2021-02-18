// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtrnáctá část
//     Programovací jazyk Go a počítačová grafika (úvod)
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-pocitacova-grafika-uvod/
//
// Demonstrační příklad číslo 6:
//     Korektní varianta předchozího příkladu

package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

const width = 256
const height = 256

func main() {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	outfile, err := os.Create("06.png")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	for x := 0; x < width; x++ {
		alpha := byte(x)
		c := color.NRGBA{0, 255, 0, alpha}
		for y := 0; y < height; y++ {
			img.SetNRGBA(x, y, c)
		}
	}
	png.Encode(outfile, img)
}
