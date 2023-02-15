// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtrnáctá část
//     Programovací jazyk Go a počítačová grafika (úvod)
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-pocitacova-grafika-uvod/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze čtrnácté části:
//    https://github.com/tisnik/go-root/blob/master/article_14/README.md
//
// Demonstrační příklad číslo 9:
//     Přímý přístup k jednotlivým pixelům; první varianta
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_14/09_raw_pixels.html

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

	outfile, err := os.Create("09.png")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	for y := 0; y < height; y++ {
		index := img.Stride * y
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
