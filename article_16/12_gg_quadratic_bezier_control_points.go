// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestnáctá část
//     Programovací jazyk Go a grafika: další užitečné funkce poskytované knihovnou GG
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-grafika-dalsi-uzitecne-funkce-poskytovane-knihovnou-gg/
//
// Demonstrační příklad číslo 12:
//     Kvadratická Bézierova křivka s řídicími body.

package main

import "github.com/fogleman/gg"

func drawQuadraticBezier(dc *gg.Context, x0 float64, y0 float64, x1 float64, y1 float64, x2 float64, y2 float64) {
	dc.SetRGBA(1.0, 0.5, 0.5, 1.0)
	dc.DrawLine(x0, y0, x1, y1)
	dc.Stroke()
	dc.DrawLine(x1, y1, x2, y2)
	dc.Stroke()

	dc.SetRGBA(0.0, 0.0, 0.0, 1.0)
	dc.MoveTo(x0, y0)
	dc.QuadraticTo(x1, y1, x2, y2)
	dc.Stroke()

	dc.SetRGBA(0.2, 0.2, 1.0, 1.0)

	dc.DrawCircle(x0, y0, 3)
	dc.Stroke()

	dc.DrawCircle(x1, y1, 3)
	dc.Stroke()

	dc.DrawCircle(x2, y2, 3)
	dc.Stroke()
}

func main() {
	const width = 320
	const height = 240

	dc := gg.NewContext(width, height)

	dc.DrawRectangle(0, 0, width, height)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()

	drawQuadraticBezier(dc, 10, 150, 50, 10, 90, 150)
	drawQuadraticBezier(dc, 110, 100, 190, 100, 150, 190)
	drawQuadraticBezier(dc, 250, 150, 210, 60, 290, 150)

	dc.SavePNG("12.png")
}
