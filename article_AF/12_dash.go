// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Demonstrační příklad číslo 12:
//    Balíček gogpu/gg.
//    Nastavení stylu vykreslovaných úseček.
//

package main

import "github.com/gogpu/gg"

func main() {
	// rozměry rastrového obrázku
	const width = 320
	const height = 240

	// jméno souboru s výsledným obrázkem
	const fileName = "12.png"

	// vytvoření kontextu
	dc := gg.NewContext(width, height)

	// pozadí obrázku
	dc.DrawRectangle(0, 0, width, height)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()

	dc.SetRGBA(0.0, 0.0, 0.0, 1)

	dc.DrawLine(32, 20, 288, 20)
	dc.Stroke()

	// styl vykreslení úsečky
	dc.SetDash(10)
	dc.DrawLine(32, 40, 288, 40)
	dc.Stroke()

	// styl vykreslení úsečky
	dc.SetDash(10, 10)
	dc.DrawLine(32, 60, 288, 60)
	dc.Stroke()

	// styl vykreslení úsečky
	dc.SetDash(10, 5)
	dc.DrawLine(32, 80, 288, 80)
	dc.Stroke()

	// styl vykreslení úsečky
	dc.SetDash(10, 5, 2, 5)
	dc.DrawLine(32, 100, 288, 100)
	dc.Stroke()

	// šířka vykreslení úsečky
	dc.SetLineWidth(4.0)
	dc.SetLineCap(gg.LineCapButt)

	// styl vykreslení úsečky
	dc.DrawLine(32, 140, 288, 140)
	dc.Stroke()

	// styl vykreslení úsečky
	dc.SetDash(10)
	dc.DrawLine(32, 160, 288, 160)
	dc.Stroke()

	// styl vykreslení úsečky
	dc.SetDash(10, 10)
	dc.DrawLine(32, 180, 288, 180)
	dc.Stroke()

	// styl vykreslení úsečky
	dc.SetDash(10, 5)
	dc.DrawLine(32, 200, 288, 200)
	dc.Stroke()

	// styl vykreslení úsečky
	dc.SetDash(10, 5, 2, 5)
	dc.DrawLine(32, 220, 288, 220)
	dc.Stroke()

	// uložení výsledného rastrového obrázku
	dc.SavePNG(fileName)
}
