// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Patnáctá část
//     Programovací jazyk Go a grafika: tvorba animovaných GIFů, grafická knihovna GG
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-grafika-tvorba-animovanych-gifu-graficka-knihovna-gg/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z patnácté části:
//    https://github.com/tisnik/go-root/blob/master/article_15/README.md
//
// Demonstrační příklad číslo 13:
//     Animace progress baru vytvořená jako animovaný GIF
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_15/13_progress_bar.html

package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"os"
)

func CreateImage(width int, height int, imageIndex int) *image.Paletted {
	var palette = []color.Color{
		color.RGBA{50, 50, 50, 255},
		color.RGBA{0, 200, 200, 255},
	}

	if imageIndex == 0 {
		img := image.NewPaletted(image.Rect(0, 0, width, height), palette)

		r := image.Rect(0, 0, width, height)
		draw.Draw(img, r, &image.Uniform{palette[0]}, image.ZP, draw.Src)

		return img
	}
	img := image.NewPaletted(image.Rect(imageIndex, 0, imageIndex+1, height), palette)

	r := image.Rect(imageIndex, 0, imageIndex+1, height)
	draw.Draw(img, r, &image.Uniform{palette[1]}, image.ZP, draw.Src)

	return img
}

func main() {
	var images []*image.Paletted
	var delays []int
	var disposals []byte

	for i := 0; i < 200; i++ {
		img := CreateImage(200, 10, i)
		images = append(images, img)
		delays = append(delays, 2)
		disposals = append(disposals, gif.DisposalNone)
	}

	outfile, err := os.Create("13.gif")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	gif.EncodeAll(outfile, &gif.GIF{
		Image:    images,
		Delay:    delays,
		Disposal: disposals,
	})
}
