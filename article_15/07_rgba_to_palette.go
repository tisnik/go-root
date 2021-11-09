// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Patnáctá část
//     Programovací jazyk Go a grafika: tvorba animovaných GIFů, grafická knihovna GG
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-grafika-tvorba-animovanych-gifu-graficka-knihovna-gg/
//
// Demonstrační příklad číslo 7:
//     Převod obrázku z prostoru RGBA do prostoru s paletou se snižováním počtu barev
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_15/07_rgba_to_palette.html

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"
)

const width = 512
const height = 512

func main() {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < 256; x++ {
		for y := 0; y < 256; y++ {
			c := color.NRGBA{0, byte(x), byte(y), 255}
			img.SetNRGBA(x, y, c)

			c = color.NRGBA{85, byte(x), byte(y), 255}
			img.SetNRGBA(x+256, y, c)

			c = color.NRGBA{170, byte(x), byte(y), 255}
			img.SetNRGBA(x, y+256, c)

			c = color.NRGBA{255, byte(x), byte(y), 255}
			img.SetNRGBA(x+256, y+256, c)
		}
	}

	for colors := 1; colors <= 256; colors <<= 1 {
		filename := fmt.Sprintf("07_%03d.gif", colors)

		outfile, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		defer outfile.Close()

		options := gif.Options{NumColors: colors, Quantizer: nil, Drawer: nil}

		gif.Encode(outfile, img, &options)
	}
}
