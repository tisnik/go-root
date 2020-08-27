// Seriál "Programovací jazyk Go"
//
// Patnáctá část
//     Programovací jazyk Go a grafika: tvorba animovaných GIFů, grafická knihovna GG
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-grafika-tvorba-animovanych-gifu-graficka-knihovna-gg/
//
// Demonstrační příklad číslo 3:
//     Export obrázku o rozlišení 1×1 pixel s 256 barvovou paletou

package main

import (
	"image"
	"image/color"
	"image/gif"
	"os"
)

// CreateImage function creates a new RGBA image with resolution width x height
func CreateImage(width int, height int) *image.RGBA {
	c := color.RGBA{255, 0, 0, 255}
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	img.Set(0, 0, c)

	return img
}

func main() {
	img := CreateImage(1, 1)

	outfile, err := os.Create("03.gif")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	err = gif.Encode(outfile, img, nil)
	if err != nil {
		panic(err)
	}
}
