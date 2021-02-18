// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Patnáctá část
//     Programovací jazyk Go a grafika: tvorba animovaných GIFů, grafická knihovna GG
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-grafika-tvorba-animovanych-gifu-graficka-knihovna-gg/
//
// Demonstrační příklad číslo 8:
//     Jednoduchá animace se dvěma snímky o velikosti 32×32 pixelů

package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"os"
)

const BoardSize = 8

func CreateImage(width int, height int, paletteIndex int) *image.Paletted {
	var palette []color.Color

	if paletteIndex == 0 {
		palette = []color.Color{
			color.RGBA{150, 150, 150, 255},
		}
	} else {
		palette = []color.Color{
			color.RGBA{250, 0, 0, 255},
		}
	}

	img := image.NewPaletted(image.Rect(0, 0, width, height), palette)

	r := image.Rect(0, 0, width, height)
	draw.Draw(img, r, &image.Uniform{palette[0]}, image.ZP, draw.Src)

	return img
}

func main() {
	images := []*image.Paletted{
		CreateImage(32, 32, 0),
		CreateImage(32, 32, 1),
	}

	delays := []int{100, 100}

	outfile, err := os.Create("08.gif")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	gif.EncodeAll(outfile, &gif.GIF{
		Image: images,
		Delay: delays,
	})
}
