// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
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

// BoardSize represents size of chessboard being rendered
const BoardSize = 8

// CreateChessboard function draws chessboard onto the test image
func CreateChessboard(width int, height int, boardSize int) *image.RGBA {
	var palette = []color.Color{
		color.RGBA{150, 205, 50, 255},
		color.RGBA{0, 100, 0, 255},
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	indexColor := 0
	horizontalBlock := int(width / boardSize)
	verticalBlock := int(height / boardSize)

	xFrom := 0
	xTo := horizontalBlock
	for x := 0; x < boardSize; x++ {
		yFrom := 0
		yTo := verticalBlock
		for y := 0; y < boardSize; y++ {
			r := image.Rect(xFrom, yFrom, xTo, yTo)
			draw.Draw(img, r, &image.Uniform{palette[indexColor]}, image.ZP, draw.Src)
			yFrom = yTo
			yTo += verticalBlock
			indexColor = 1 - indexColor
		}
		xFrom = xTo
		xTo += horizontalBlock
		indexColor = 1 - indexColor
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
