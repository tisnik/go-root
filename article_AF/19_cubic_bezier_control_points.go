// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Demonstrační příklad číslo 19:
//    Balíček gogpu/gg.
//    Kubická Bézierova křivka s řídicími body.
//

package main

import "github.com/gogpu/gg"

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
	// vykreslení
	dc.Stroke()

	dc.DrawCircle(x2, y2, 3)
	// vykreslení
	dc.Stroke()

	dc.DrawCircle(x3, y3, 3)
	// vykreslení
	dc.Stroke()
}

func clearCanvas(dc *gg.Context, width float64, height float64) {
	dc.DrawRectangle(0, 0, width, height)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()
}

func main() {
	// rozměry rastrového obrázku
	const width = 320
	const height = 240

	// jméno souboru s výsledným obrázkem
	const fileName1 = "19_1.png"
	const fileName2 = "19_2.png"

	// vytvoření kontextu
	dc := gg.NewContext(width, height)
	// vymazání obrázku
	clearCanvas(dc, width, height)

	// vykreslení cest
	drawCubicBezier(dc, 10, 180, 10, 10, 120, 180, 120, 10)
	drawCubicBezier(dc, 230, 180, 280, 60, 230, 60, 280, 180)

	// uložení výsledného rastrového obrázku
	dc.SavePNG(fileName1)

	// vymazání obrázku
	clearCanvas(dc, width, height)

	// vykreslení cesty
	drawCubicBezier(dc, 130, 180, 210, 100, 100, 100, 180, 180)

	// uložení výsledného rastrového obrázku
	dc.SavePNG(fileName2)
}
