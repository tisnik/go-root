// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Demonstrační příklad číslo 5:
//    Balíček gogpu/gg.
//    Vymazání obrazovky (varianta se specifikací barvy).
//

package main

import "github.com/gogpu/gg"

func main() {
	// rozměry rastrového obrázku
	const width = 320
	const height = 240

	// jméno souboru s výsledným obrázkem
	const fileName = "05.png"

	// vytvoření kontextu
	dc := gg.NewContext(width, height)

	// vykreslování s využitím kontextu
	dc.ClearWithColor(gg.RGBA{R: 0.5, G: 0.9, B: 0.0, A: 0.5})

	// uložení výsledného rastrového obrázku
	dc.SavePNG(fileName)
}
