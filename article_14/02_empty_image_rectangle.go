// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtrnáctá část
//     Programovací jazyk Go a počítačová grafika (úvod)
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-pocitacova-grafika-uvod/
//
// Seznam příkladů ze čtrnácté části:
//    https://github.com/tisnik/go-root/blob/master/article_14/README.md
//
// Demonstrační příklad číslo 2:
//     Alternativní způsob vytvoření obdélníku pro určení rozměrů obrázku
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_14/02_empty_image_rectangle.html

package main

import (
	"image"
	"image/png"
	"os"
)

const width = 256
const height = 256

func main() {
	rect := image.Rectangle{image.Point{0, 0},
		image.Point{width, height}}

	img := image.NewNRGBA(rect)

	outfile, err := os.Create("02.png")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	png.Encode(outfile, img)
}
