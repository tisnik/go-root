// Seriál "Programovací jazyk Go"
//
// Čtrnáctá část
//     Programovací jazyk Go a počítačová grafika (úvod)
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-pocitacova-grafika-uvod/
//
// Demonstrační příklad číslo 14:
//     Základní vlastnosti barvového prostoru RGBA

package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

const width = 512
const height = 512

func main() {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	outfile, err := os.Create("14.png")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

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
	png.Encode(outfile, img)
}
