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
// Demonstrační příklad číslo 9:
//     Složitější animace pohybující se šachovnice
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_15/09_gif_animation.html

package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"os"
)

const BoardSize = 8

func CreateChessboard(width int, height int, board_size int, xoffset int, yoffset int, step_size int) *image.Paletted {
	var palette = []color.Color{
		color.RGBA{150, 205, 50, 255},
		color.RGBA{0, 100, 0, 255},
	}

	img := image.NewPaletted(image.Rect(0, 0, width, height), palette)

	indexColor := 0
	horBlock := int(width / board_size)
	verBlock := int(height / board_size)

	xFrom := 0
	xTo := horBlock
	for x := 0; x < board_size+step_size; x++ {
		yFrom := 0
		yTo := verBlock
		for y := 0; y < board_size+step_size; y++ {
			r := image.Rect(xFrom+xoffset, yFrom+yoffset, xTo+xoffset, yTo+yoffset)
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
	var images []*image.Paletted
	var delays []int

	steps := 2 * 256 / BoardSize
	for step := 0; step < steps; step++ {
		img := CreateChessboard(256, 256, BoardSize, step*2-steps*2, step-steps, steps*2)
		images = append(images, img)
		delays = append(delays, 10)
	}

	outfile, err := os.Create("09.gif")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	gif.EncodeAll(outfile, &gif.GIF{
		Image: images,
		Delay: delays,
	})
}
