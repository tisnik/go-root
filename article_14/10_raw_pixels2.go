// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtrnáctá část
//     Programovací jazyk Go a počítačová grafika (úvod)
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-pocitacova-grafika-uvod/
//
// Demonstrační příklad číslo 10:
//     Přímý přístup k jednotlivým pixelům; druhá varianta
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_14/10_raw_pixels2.html

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

	outfile, err := os.Create("10.png")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	for y := 0; y < height; y++ {
		scanline := img.Pix[img.Stride*y:]
		i := 0
		for x := 0; x < width; x++ {
			scanline[i] = 0
			i++
			scanline[i] = byte(y)
			i++
			scanline[i] = 255
			i++
			scanline[i] = byte(x)
			i++
		}
	}
	png.Encode(outfile, img)
}
