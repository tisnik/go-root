// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Demonstrační příklad číslo 6:
//    Balíček gogpu/gg.
//    Práce s barvami.
//

package main

import "github.com/gogpu/gg"

func main() {
	// rozměry rastrového obrázku
	const width = 256
	const height = 256

	// jméno souboru s výsledným obrázkem
	const fileName = "06.png"

	// vytvoření kontextu
	dc := gg.NewContext(width, height)

	for y := range 16 {
		for x := range 16 {
			// nastavení barvy kreslení
			dc.SetRGB(float64(x)/16.0, 0.5, float64(y)/16.0)
			// vykreslení vyplněného obdélníka
			dc.DrawRectangle(float64(x*16)+3, float64(y*16)+3, 10, 10)
			dc.Fill()
		}
	}

	// uložení výsledného rastrového obrázku
	dc.SavePNG(fileName)
}
