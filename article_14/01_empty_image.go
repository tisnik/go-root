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
// Demonstrační příklad číslo 1:
//     Vytvoření rastrového obrázku s výchozími hodnotami pixelů (zcela průhledné černé pixely)
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_14/01_empty_image.html

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
