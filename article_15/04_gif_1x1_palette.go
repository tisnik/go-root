// Seriál "Programovací jazyk Go"
//
// Patnáctá část
//     Programovací jazyk Go a grafika: tvorba animovaných GIFů, grafická knihovna GG
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-grafika-tvorba-animovanych-gifu-graficka-knihovna-gg/
//
// Demonstrační příklad číslo 4:
//     Export obrázku o rozlišení 1×1 pixel s dvoubarevnou paletou

package main

import (
	"image"
	"image/color"
	"image/gif"
	"os"
)

// CreateImage function creates a new image with color palette and with
// resolution width x height
func CreateImage(width int, height int) *image.Paletted {
	var palette = []color.Color{
		color.RGBA{255, 128, 128, 255},
		color.RGBA{255, 255, 255, 255},
	}

	c := color.RGBA{255, 128, 128, 255}
	img := image.NewPaletted(image.Rect(0, 0, width, height), palette)
	img.Set(0, 0, c)

	return img
}

func main() {
	img := CreateImage(1, 1)

	outfile, err := os.Create("04.gif")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	err = gif.Encode(outfile, img, nil)
	if err != nil {
		panic(err)
	}
}
