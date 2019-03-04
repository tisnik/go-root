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

	index_color := 0
	board_size := 8
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
	png.Encode(outfile, img)
}
