// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Patnáctá část
//     Programovací jazyk Go a grafika: tvorba animovaných GIFů, grafická knihovna GG
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-grafika-tvorba-animovanych-gifu-graficka-knihovna-gg/
//
// Seznam příkladů z patnácté části:
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

func CreateChessboard(width int, height int, board_size int) *image.Paletted {
	var palette = []color.Color{
		color.RGBA{150, 205, 50, 255},
		color.RGBA{0, 100, 0, 255},
	}

	img := image.NewPaletted(image.Rect(0, 0, width, height), palette)

	index_color := 0
	hor_block := int(width / board_size)
	ver_block := int(height / board_size)

	x_from := 0
	x_to := hor_block
	for x := 0; x < board_size; x++ {
		y_from := 0
		y_to := ver_block
		for y := 0; y < board_size; y++ {
			r := image.Rect(x_from, y_from, x_to, y_to)
			draw.Draw(img, r, &image.Uniform{palette[index_color]}, image.ZP, draw.Src)
			y_from = y_to
			y_to += ver_block
			index_color = 1 - index_color
		}
		x_from = x_to
		x_to += hor_block
		index_color = 1 - index_color
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
