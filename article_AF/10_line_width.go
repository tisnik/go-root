// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Demonstrační příklad číslo 9:
//    Balíček gogpu/gg.
//    Nastavení šířky vykreslovaných úseček a křivek.
//

package main

import "github.com/gogpu/gg"

func main() {
	// rozměry rastrového obrázku
	const width = 320
	const height = 240

	// jméno souboru s výsledným obrázkem
	const fileName = "10.png"

	// vytvoření kontextu
	dc := gg.NewContext(width, height)

	// pozadí obrázku
	dc.DrawRectangle(0, 0, width, height)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()

	dc.SetRGBA(0.0, 0.0, 0.0, 1)

	for i := 0; i < 256; i += 16 {
		// nastavení šířky úsečky
		width := float64(i) / 20
		dc.SetLineWidth(width)

		x := float64(i + 32)
		dc.DrawLine(x, 20, x, height-20)

		// vykreslení
		dc.Stroke()
	}

	// uložení výsledného rastrového obrázku
	dc.SavePNG(fileName)
}
