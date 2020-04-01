// Seriál "Programovací jazyk Go"
//
// Čtrnáctá část
//     Programovací jazyk Go a počítačová grafika (úvod)
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-pocitacova-grafika-uvod/
//
// Demonstrační příklad číslo 11:
//     Přímý přístup k jednotlivým pixelům; třetí varianta

package main

import (
	"image"
	"image/png"
	"os"
)

const width = 256
const height = 256

func main() {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	outfile, err := os.Create("11.png")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	for y := 0; y < height; y++ {
		index := img.PixOffset(0, y)
		for x := 0; x < width; x++ {
			img.Pix[index] = 0
			index++
			img.Pix[index] = 0
			index++
			img.Pix[index] = 255
			index++
			img.Pix[index] = byte(x)
			index++
		}
	}
	png.Encode(outfile, img)
}
