// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Demonstrační příklad číslo 15:
//    Balíček gogpu/gg.
//    Uzavřená cesta (closed path).

package main

import "github.com/gogpu/gg"

func main() {
	// rozměry rastrového obrázku
	const width = 320
	const height = 240

	// jméno souboru s výsledným obrázkem
	const fileName = "15.png"

	// vytvoření kontextu
	dc := gg.NewContext(width, height)

	// pozadí obrázku
	dc.DrawRectangle(0, 0, width, height)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()

	dc.SetRGBA(0.0, 0.0, 0.0, 1)

	// začátek vykreslování cesty
	dc.MoveTo(100, 200)
	dc.LineTo(100, 100)
	dc.LineTo(150, 50)
	dc.LineTo(200, 100)
	dc.LineTo(200, 200)
	dc.ClosePath()

	// vlastní vykreslení cesty
	dc.Stroke()

	// uložení výsledného rastrového obrázku
	dc.SavePNG(fileName)
}
