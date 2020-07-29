// Seriál "Programovací jazyk Go"
//
// Patnáctá část
//     Programovací jazyk Go a grafika: tvorba animovaných GIFů, grafická knihovna GG
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-grafika-tvorba-animovanych-gifu-graficka-knihovna-gg/
//
// Demonstrační příklad číslo 1:
//     Ukázka blendingu

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

func DrawHorizontalLine(img *image.NRGBA, color color.Color, x1 int, x2 int, y int) {
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	for x := x1; x < x2; x++ {
		img.Set(x, y, color)
	}
}

func DrawVerticalLine(img *image.NRGBA, color color.Color, x int, y1 int, y2 int) {
	if y1 > y2 {
		y1, y2 = y2, y1
	}
	for y := y1; y < y2; y++ {
		img.Set(x, y, color)
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Step(v1 int, v2 int) int {
	if v1 < v2 {
		return 1
	} else {
		return -1
	}
}

func DrawLine(img *image.NRGBA, color color.Color, x1 int, y1 int, x2 int, y2 int) {
	// specialni pripad - svisla usecka
	if x1 == x2 {
		DrawVerticalLine(img, color, x1, y1, y2)
		return
	}

	// specialni pripad - vodorovna usecka
	if y1 == y2 {
		DrawHorizontalLine(img, color, x1, x2, y1)
		return
	}

	// takze mame smulu a musime pouzit plnou verzi algoritmu

	// zrcadleni algoritmu pro dalsi oktanty
	x := x1
	y := y1

	// konstanty pouzite pri vykreslovani
	dx := Abs(x2 - x1)
	dy := Abs(y2 - y1)
	sx := Step(x1, x2)
	sy := Step(y1, y2)

	// pocatecni hodnota akumulatoru chyby
	err := dx >> 1
	if dx <= dy {
		err = -dy >> 1
	}

	// vse je pripraveno k vlastnimu vykresleni usecky
	for {
		img.Set(x, y, color)
		// test, zda se jiz doslo k poslednimu bodu
		if x == x2 && y == y2 {
			break
		}
		e2 := err
		if e2 > -dx {
			// prepocet kumulovane chyby
			err -= dy
			// posun na predchozi ci dalsi pixel na radku
			x += sx
		}
		if e2 < dy {
			// prepocet kumulovane chyby
			err += dx
			// posun na predchozi ci nasledujici radek
			y += sy
		}
	}
}

func CreateStringArt(width int, height int) draw.Image {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	c := color.RGBA{255, 255, 255, 255}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, c)
		}
	}

	c = color.RGBA{0, 0, 255, 255}
	DrawLine(img, c, 20, 10, 245, 10)

	c = color.RGBA{255, 0, 0, 255}
	DrawLine(img, c, 10, 20, 10, 245)

	c = color.RGBA{0, 255, 0, 255}
	DrawLine(img, c, 20, 20, width>>1, height>>1)

	c = color.RGBA{0, 0, 0, 255}
	for x := 10; x < width; x += 10 {
		DrawLine(img, c, width-5, x, width-x, height-5)
	}

	return img
}

func CreateChessboard(width int, height int) draw.Image {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	palette := make(map[int]color.RGBA, 2)

	palette[0] = color.RGBA{150, 205, 50, 255}
	palette[1] = color.RGBA{0, 100, 0, 0}

	index_color := 0
	board_size := 8
	horizontalBlock := int(width / board_size)
	verticalBlock := int(height / board_size)

	xFrom := 0
	x_to := horizontalBlock
	for x := 0; x < board_size; x++ {
		y_from := 0
		y_to := verticalBlock
		for y := 0; y < board_size; y++ {
			r := image.Rect(xFrom, y_from, x_to, y_to)
			draw.Draw(img, r, &image.Uniform{palette[index_color]}, image.ZP, draw.Src)
			y_from = y_to
			y_to += verticalBlock
			index_color = 1 - index_color
		}
		xFrom = x_to
		x_to += horizontalBlock
		index_color = 1 - index_color
	}
	return img
}

func main() {
	img1 := CreateChessboard(width, height)
	img2 := CreateStringArt(width, height)

	draw.Draw(img2, img2.Bounds(), img1, image.ZP, draw.Over)

	outfile, err := os.Create("01.png")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	png.Encode(outfile, img2)
}
