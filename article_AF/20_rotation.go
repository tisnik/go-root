// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sto třináctá část
//    Tvorba 2D i 3D grafiky a animací v Go s využitím projektu GoGPU
//    https://www.root.cz/clanky/tvorba-2d-i-3d-grafiky-a-animaci-v-go-s-vyuzitim-projektu-gogpu/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze sto třinácté části:
//    https://github.com/tisnik/go-root/blob/master/article_AF/README.md
//
// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Demonstrační příklad číslo 20:
//    Balíček gogpu/gg.
//    Rotace vykreslovaných objektů.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_AF/20_rotation.html

package main

import (
	"fmt"
	"math"

	"github.com/gogpu/gg"
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

func Radians(angle float64) float64 {
	return float64(angle / 180.0 * math.Pi)
}

func main() {
	// rozměry rastrového obrázku
	const width = 320
	const height = 240

	// vytvoření kontextu
	dc := gg.NewContext(width, height)

	for alpha := 0; alpha < 360; alpha += 30 {
		clearCanvas(dc, width, height)
		// transformace a vykreslení
		dc.Push()
		dc.RotateAbout(Radians(float64(alpha)), width/2, height/2)
		drawCubicBezier(dc, 130, 180, 210, 100, 100, 100, 180, 180)
		dc.Pop()
		filename := fmt.Sprintf("20_%03d.png", alpha)
		dc.SavePNG(filename)
	}

}
