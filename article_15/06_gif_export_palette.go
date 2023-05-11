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
// Demonstrační příklad číslo 6:
//     Šachovnice vykreslená s dvoubarevnou paletou
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_15/06_gif_export_palette.html

package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"os"
)

const BoardSize = 8

func CreateChessboard(width int, height int, boardSize int) *image.Paletted {
	var palette = []color.Color{
		color.RGBA{150, 205, 50, 255},
		color.RGBA{0, 100, 0, 255},
	}

	img := image.NewPaletted(image.Rect(0, 0, width, height), palette)

	indexColor := 0
	horBlock := int(width / boardSize)
	verBlock := int(height / boardSize)

	xFrom := 0
	xTo := horBlock
	for x := 0; x < boardSize; x++ {
		yFrom := 0
		yTo := verBlock
		for y := 0; y < boardSize; y++ {
			r := image.Rect(xFrom, yFrom, xTo, yTo)
			draw.Draw(img, r, &image.Uniform{palette[indexColor]}, image.ZP, draw.Src)
			yFrom = yTo
			yTo += verBlock
			indexColor = 1 - indexColor
		}
		xFrom = xTo
		xTo += horBlock
		indexColor = 1 - indexColor
	}
	return img
}

func main() {
	img := CreateChessboard(256, 256, BoardSize)

	outfile, err := os.Create("06.gif")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	err = gif.Encode(outfile, img, nil)
	if err != nil {
		panic(err)
	}
}
