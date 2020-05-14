// Seriál "Programovací jazyk Go"
//
// Čtrnáctá část
//     Programovací jazyk Go a počítačová grafika (úvod)
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-pocitacova-grafika-uvod/
//
// Demonstrační příklad číslo 19:
//     Blending v Go (opět rastrové operace)

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

// DrawHorizontalLine draws horizontal line from [x1,y] to [x2,y] by specified color
func DrawHorizontalLine(img *image.NRGBA, color color.Color, x1 int, x2 int, y int) {
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	for x := x1; x < x2; x++ {
		img.Set(x, y, color)
	}
}

// DrawVerticalLine draws vertical line from [x,y1] to [x,y2] by specified color
func DrawVerticalLine(img *image.NRGBA, color color.Color, x int, y1 int, y2 int) {
	if y1 > y2 {
		y1, y2 = y2, y1
	}
	for y := y1; y < y2; y++ {
		img.Set(x, y, color)
	}
}

// Abs computes absolute value for given signed integer
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Step computes whether the horizontal/vertical step should be positive or negative
func Step(v1 int, v2 int) int {
	// is first coordinate less than the second one?
	if v1 < v2 {
		// if yes, step up or right
		return 1
	}
	// if no, step down or left
	return -1
}

// DrawLine draws line from [x1,y1] to [x2,y2] by specified color
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
	return img
}

func main() {
	img1 := CreateChessboard(width, height)
	img2 := CreateStringArt(width, height)

	draw.Draw(img2, img2.Bounds(), img1, image.ZP, draw.Over)

	outfile, err := os.Create("19.png")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	png.Encode(outfile, img2)
}
