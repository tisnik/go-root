// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtrnáctá část
//     Programovací jazyk Go a počítačová grafika (úvod)
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-pocitacova-grafika-uvod/
//
// Demonstrační příklad číslo 3:
//     Rastrový obrázek vyplněný konstantní barvou, použití metody SetRGBA
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_14/03_filled_image_setrgba.html

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
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	outfile, err := os.Create("03.png")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	c := color.RGBA{0, 255, 0, 255}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.SetRGBA(x, y, c)
		}
	}
	png.Encode(outfile, img)
}
