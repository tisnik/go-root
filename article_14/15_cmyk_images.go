// Seriál "Programovací jazyk Go"
//
// Čtrnáctá část
//     Programovací jazyk Go a počítačová grafika (úvod)
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-pocitacova-grafika-uvod/
//
// Demonstrační příklad číslo 15:
//     Základní vlastnosti barvového prostoru CMYK

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
	img := image.NewCMYK(image.Rect(0, 0, width, height))

	outfile, err := os.Create("15.png")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	for x := 0; x < 256; x++ {
		for y := 0; y < 256; y++ {
			c := color.CMYK{byte(x), 0, 0, 0}
			img.SetCMYK(x, y, c)

			c = color.CMYK{0, byte(x), 0, 0}
			img.SetCMYK(x+256, y, c)

			c = color.CMYK{0, 0, byte(x), 0}
			img.SetCMYK(x, y+256, c)

			c = color.CMYK{0, 0, 0, byte(x)}
			img.SetCMYK(x+256, y+256, c)
		}
	}
	png.Encode(outfile, img)
}
