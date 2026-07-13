// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Demonstrační příklad číslo 16:
//    Balíček gogpu/gg.
//    Kvadratická Bézierova křivka.
//

package main

import "github.com/gogpu/gg"

func main() {
	// rozměry rastrového obrázku
	const width = 320
	const height = 240

	// jméno souboru s výsledným obrázkem
	const fileName = "16.png"

	// vytvoření kontextu
	dc := gg.NewContext(width, height)

	// pozadí obrázku
	dc.DrawRectangle(0, 0, width, height)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()

	// cesta
	dc.SetRGBA(1.0, 0.0, 0.0, 1.0)
	dc.MoveTo(10, 150)
	dc.QuadraticTo(50, 10, 90, 150)
	// vlastní vykreslení cesty
	dc.Stroke()

	// cesta
	dc.SetRGBA(0.0, 1.0, 0.0, 1.0)
	dc.MoveTo(110, 100)
	dc.QuadraticTo(190, 100, 150, 190)
	// vlastní vykreslení cesty
	dc.Stroke()

	// cesta
	dc.SetRGBA(0.0, 0.0, 1.0, 1.0)
	dc.MoveTo(250, 150)
	dc.QuadraticTo(210, 60, 290, 150)
	// vlastní vykreslení cesty
	dc.Stroke()

	// uložení výsledného rastrového obrázku
	dc.SavePNG(fileName)
}
