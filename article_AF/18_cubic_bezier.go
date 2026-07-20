// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Demonstrační příklad číslo 18:
//    Balíček gogpu/gg.
//    Kubická Bézierova křivka.
//

package main

import "github.com/gogpu/gg"

func main() {
	// rozměry rastrového obrázku
	const width = 320
	const height = 240

	// jméno souboru s výsledným obrázkem
	const fileName = "18.png"

	// vytvoření kontextu
	dc := gg.NewContext(width, height)

	// pozadí obrázku
	dc.DrawRectangle(0, 0, width, height)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()

	// cesta
	dc.SetRGBA(1.0, 0.0, 0.0, 1.0)
	dc.MoveTo(10, 180)
	dc.CubicTo(10, 10, 120, 180, 120, 10)
	// vlastní vykreslení cesty
	dc.Stroke()

	// cesta
	dc.SetRGBA(0.0, 1.0, 0.0, 1.0)
	dc.MoveTo(110, 180)
	dc.CubicTo(190, 100, 80, 100, 160, 180)
	// vlastní vykreslení cesty
	dc.Stroke()

	// cesta
	dc.SetRGBA(0.0, 0.0, 1.0, 1.0)
	dc.MoveTo(230, 180)
	dc.CubicTo(280, 60, 230, 60, 280, 180)
	// vlastní vykreslení cesty
	dc.Stroke()

	// uložení výsledného rastrového obrázku
	dc.SavePNG(fileName)
}
