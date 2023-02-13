// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třináctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (práce s JSONem a rastrovými obrázky)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-prace-s-jsonem-a-rastrovymi-obrazky/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze třinácté části:
//    https://github.com/tisnik/go-root/blob/master/article_13/README.md
//
// Demonstrační příklad číslo 21:
//     Export rastrového obrázku do formátu JPEG s ovlivněním kvality
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_13/21_jpeg_output_low_quality.html

package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

const width = 256
const height = 256

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			var red uint8 = uint8(x)
			var green uint8 = uint8((x - y))
			var blue uint8 = uint8(y)
			c := color.RGBA{red, green, blue, 255}
			img.SetRGBA(x, y, c)
		}
	}

	outfile, err := os.Create("test2.jpeg")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	jpeg.Encode(outfile, img, &jpeg.Options{Quality: 1})
}
