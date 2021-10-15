// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtrnáctá část
//     Programovací jazyk Go a počítačová grafika (úvod)
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-pocitacova-grafika-uvod/
//
// Demonstrační příklad číslo 8:
//     Interní struktura záznamu s informacemi o obrázku
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_14/08_image_internals.html

package main

import (
	"image"
)

const width = 256
const height = 256

func main() {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	println("Stride: ", img.Stride)
	println("[]byte: ", len(img.Pix))
	r := img.Rect
	println("Rectangle:")
	println("    point 1: ", r.Min.X, r.Min.Y)
	println("    point 2: ", r.Max.X, r.Max.Y)
}
