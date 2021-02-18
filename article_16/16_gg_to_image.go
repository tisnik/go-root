// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestnáctá část
//     Programovací jazyk Go a grafika: další užitečné funkce poskytované knihovnou GG
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-grafika-dalsi-uzitecne-funkce-poskytovane-knihovnou-gg/
//
// Demonstrační příklad číslo 16:
//     Rasterizace do zvoleného rastrového formátu.

package main

import (
	"github.com/fogleman/gg"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"os"
)

func drawCubicBezier(dc *gg.Context,
	x0 float64, y0 float64, x1 float64, y1 float64,
	x2 float64, y2 float64, x3 float64, y3 float64) {

	dc.SetRGBA(1.0, 0.5, 0.5, 1.0)
	dc.DrawLine(x0, y0, x1, y1)
	dc.Stroke()
	dc.DrawLine(x1, y1, x2, y2)
	dc.Stroke()
	dc.DrawLine(x2, y2, x3, y3)
	dc.Stroke()

	dc.SetRGBA(0.0, 0.0, 0.0, 1.0)
	dc.MoveTo(x0, y0)
	dc.CubicTo(x1, y1, x2, y2, x3, y3)
	dc.Stroke()

	dc.SetRGBA(0.2, 0.2, 1.0, 1.0)
	dc.DrawCircle(x0, y0, 3)
	dc.Stroke()

	dc.DrawCircle(x1, y1, 3)
	dc.Stroke()

	dc.DrawCircle(x2, y2, 3)
	dc.Stroke()

	dc.DrawCircle(x2, y2, 3)
	dc.Stroke()

	dc.DrawCircle(x3, y3, 3)
	dc.Stroke()
}

func clearCanvas(dc *gg.Context, width float64, height float64) {
	dc.DrawRectangle(0, 0, width, height)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()
}

func main() {
	const width = 320
	const height = 240

	dc := gg.NewContext(width, height)

	var images []*image.Paletted
	var delays []int

	for alpha := 0; alpha < 360; alpha += 5 {
		clearCanvas(dc, width, height)
		dc.Push()
		dc.RotateAbout(gg.Radians(float64(alpha)), width/2, height/2)
		drawCubicBezier(dc, 130, 180, 210, 100, 100, 100, 180, 180)
		dc.Pop()
		sourceImage := dc.Image()
		palettedImage := image.NewPaletted(image.Rect(0, 0, width, height), palette.Plan9)
		draw.Draw(palettedImage, palettedImage.Rect, sourceImage, image.ZP, draw.Over)
		images = append(images, palettedImage)
		delays = append(delays, 10)
		println(alpha, "of", 360)
	}

	outfile, err := os.Create("16.gif")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	gif.EncodeAll(outfile, &gif.GIF{
		Image: images,
		Delay: delays,
	})
}
