// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestnáctá část
//     Programovací jazyk Go a grafika: další užitečné funkce poskytované knihovnou GG
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-grafika-dalsi-uzitecne-funkce-poskytovane-knihovnou-gg/
//
// Demonstrační příklad číslo 14:
//     Kubická Bézierova křivka s řídicími body.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_16/14_gg_cubic_bezier_control_points.html

package main

import "github.com/fogleman/gg"

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
	clearCanvas(dc, width, height)

	drawCubicBezier(dc, 10, 180, 10, 10, 120, 180, 120, 10)
	drawCubicBezier(dc, 230, 180, 280, 60, 230, 60, 280, 180)

	dc.SavePNG("14A.png")

	clearCanvas(dc, width, height)

	drawCubicBezier(dc, 130, 180, 210, 100, 100, 100, 180, 180)

	dc.SavePNG("14B.png")
}
