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
// Demonstrační příklad číslo 11:
//     Nastavení metody přepínání snímků: vykreslení pozadí
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_15/11_background_disposal_methods.html

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
		color.RGBA{150, 150, 150, 0},
		color.RGBA{250, 0, 0, 255},
	}

	if imageIndex == 0 {
		img := image.NewPaletted(image.Rect(0, 0, width, height), palette)

		r := image.Rect(0, 0, width, height)
		draw.Draw(img, r, &image.Uniform{palette[0]}, image.ZP, draw.Src)

		r2 := image.Rect(0, 0, 50, height)
		draw.Draw(img, r2, &image.Uniform{palette[1]}, image.ZP, draw.Src)

		return img
	} else {
		img := image.NewPaletted(image.Rect(imageIndex*60, 0, imageIndex*60+50, height), palette)

		r2 := image.Rect(imageIndex*60, 0, imageIndex*60+50, height)
		draw.Draw(img, r2, &image.Uniform{palette[1]}, image.ZP, draw.Src)

		return img
	}

}

func main() {
	var images []*image.Paletted
	var delays []int
	var disposals []byte

	for i := 0; i < 6; i++ {
		img := CreateImage(350, 50, i)
		images = append(images, img)
		delays = append(delays, 100)
		disposals = append(disposals, gif.DisposalBackground)
	}

	outfile, err := os.Create("11.gif")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	gif.EncodeAll(outfile, &gif.GIF{
		Image:           images,
		Delay:           delays,
		Disposal:        disposals,
		BackgroundIndex: 0,
	})
}
