// Seriál "Programovací jazyk Go"
//
// Patnáctá část
//     Programovací jazyk Go a grafika: tvorba animovaných GIFů, grafická knihovna GG
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-grafika-tvorba-animovanych-gifu-graficka-knihovna-gg/
//
// Demonstrační příklad číslo 2:
//     Export obrázku šachovnice do formátu GIF

package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"os"
)

const BoardSize = 8

func CreateChessboard(width int, height int, board_size int) *image.RGBA {
	var palette = []color.Color{
		color.RGBA{150, 205, 50, 255},
		color.RGBA{0, 100, 0, 255},
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	index_color := 0
	horizontalBlock := int(width / board_size)
	verticalBlock := int(height / board_size)

	xFrom := 0
	xTo := horizontalBlock
	for x := 0; x < board_size; x++ {
		yFrom := 0
		yTo := verticalBlock
		for y := 0; y < board_size; y++ {
			r := image.Rect(xFrom, yFrom, xTo, yTo)
			draw.Draw(img, r, &image.Uniform{palette[index_color]}, image.ZP, draw.Src)
			yFrom = yTo
			yTo += verticalBlock
			index_color = 1 - index_color
		}
		xFrom = xTo
		xTo += horizontalBlock
		index_color = 1 - index_color
	}
	return img
}

func main() {
	img := CreateChessboard(256, 256, BoardSize)

	outfile, err := os.Create("02.gif")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	err = gif.Encode(outfile, img, nil)
	if err != nil {
		panic(err)
	}
}
