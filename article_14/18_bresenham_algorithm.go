// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtrnáctá část
//     Programovací jazyk Go a počítačová grafika (úvod)
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-pocitacova-grafika-uvod/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze čtrnácté části:
//    https://github.com/tisnik/go-root/blob/master/article_14/README.md
//
// Demonstrační příklad číslo 18:
//     Implementace Bresenhamova algoritmu pro vykreslení úsečky
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_14/18_bresenham_algorithm.html

package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

const width = 256
const height = 256

// DrawHorizontalLine function draws horizontal line from [x1, y] to [x2, y] into the given image
func DrawHorizontalLine(img *image.RGBA, lineColor color.Color, x1, x2, y int) {
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	for x := x1; x < x2; x++ {
		img.Set(x, y, lineColor)
	}
}

// DrawVerticalLine function draws vertical line from [x, y1] to [x, y2] into the given image
func DrawVerticalLine(img *image.RGBA, lineColor color.Color, x, y1, y2 int) {
	if y1 > y2 {
		y1, y2 = y2, y1
	}
	for y := y1; y < y2; y++ {
		img.Set(x, y, lineColor)
	}
}

// Abs function computer absolute value for given integer input
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Step function computes step direction (left/right, up/down) for Bresenham algorithm
func Step(v1 int, v2 int) int {
	// is first coordinate less than the second one?
	if v1 < v2 {
		// if yes, step up or right
		return 1
	}
	// if no, step down or left
	return -1
}

// DrawLine function draws line from [x1, y1] to [x2, y2] into the given image
func DrawLine(img *image.RGBA, lineColor color.Color, x1, y1, x2, y2 int) {
	// specialni pripad - svisla usecka
	if x1 == x2 {
		DrawVerticalLine(img, lineColor, x1, y1, y2)
		return
	}

	// specialni pripad - vodorovna usecka
	if y1 == y2 {
		DrawHorizontalLine(img, lineColor, x1, x2, y1)
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
		img.Set(x, y, lineColor)
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

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	outfile, err := os.Create("18.png")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

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
	png.Encode(outfile, img)
}
