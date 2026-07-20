// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Demonstrační příklad číslo 11:
//    Balíček gogpu/gg.
//    Nastavení stylu konců úseček.
//

package main

import "github.com/gogpu/gg"

func main() {
	// rozměry rastrového obrázku
	const width = 320
	const height = 240

	// jméno souboru s výsledným obrázkem
	const fileName = "11.png"

	// vytvoření kontextu
	dc := gg.NewContext(width, height)

	// pozadí obrázku
	dc.DrawRectangle(0, 0, width, height)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()

	dc.SetRGBA(0.0, 0.0, 0.0, 1.0)

	dc.SetLineWidth(10.0)

	for i := 0; i < 256; i += 16 {
		x := float64(i + 32)

		width := float64(i) / 20
		dc.SetLineWidth(width)

		// nastavení stylů obou konců úsečky
		dc.SetLineCap(gg.LineCapRound)
		dc.DrawLine(x, 20, x, 75)
		dc.Stroke()

		// nastavení stylů obou konců úsečky
		dc.SetLineCap(gg.LineCapButt)
		dc.DrawLine(x, 92, x, height-92)
		dc.Stroke()

		// nastavení stylů obou konců úsečky
		dc.SetLineCap(gg.LineCapSquare)
		dc.DrawLine(x, height-75, x, height-20)
		dc.Stroke()
	}

	dc.SetRGBA(1.0, 0.0, 0.0, 1.0)
	dc.SetLineWidth(1.0)

	dc.DrawLine(32, 20, width-32, 20)
	dc.Stroke()

	dc.DrawLine(32, 75, width-32, 75)
	dc.Stroke()

	dc.DrawLine(32, 92, width-32, 92)
	dc.Stroke()

	dc.DrawLine(32, height-92, width-32, height-92)
	dc.Stroke()

	dc.DrawLine(32, height-75, width-32, height-75)
	dc.Stroke()

	dc.DrawLine(32, height-20, width-32, height-20)
	dc.Stroke()

	// uložení výsledného rastrového obrázku
	dc.SavePNG(fileName)
}
