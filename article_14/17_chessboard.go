// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtrnáctá část
//     Programovací jazyk Go a počítačová grafika (úvod)
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-pocitacova-grafika-uvod/
//
// Demonstrační příklad číslo 17:
//     Využití balíčku draw pro vykreslení šachovnice (rastrové operace)
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_14/17_chessboard.html

package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

const width = 256
const height = 256

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	outfile, err := os.Create("17.png")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	palette := make(map[int]color.RGBA, 2)

	palette[0] = color.RGBA{150, 205, 50, 255}
	palette[1] = color.RGBA{0, 100, 0, 255}

	indexColor := 0
	boardSize := 8
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
	png.Encode(outfile, img)
}
