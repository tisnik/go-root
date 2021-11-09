// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Patnáctá část
//     Programovací jazyk Go a grafika: tvorba animovaných GIFů, grafická knihovna GG
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-grafika-tvorba-animovanych-gifu-graficka-knihovna-gg/
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

	index_color := 0
	hor_block := int(width / board_size)
	ver_block := int(height / board_size)

	x_from := 0
	x_to := hor_block
	for x := 0; x < board_size+step_size; x++ {
		y_from := 0
		y_to := ver_block
		for y := 0; y < board_size+step_size; y++ {
			r := image.Rect(x_from+xoffset, y_from+yoffset, x_to+xoffset, y_to+yoffset)
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
