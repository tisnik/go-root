// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Demonstrační příklad číslo 3:
//    Balíček gogpu/gg.
//    Vymazání obrazovky (lepší varianta).
//

package main

import "github.com/gogpu/gg"

func main() {
	// rozměry rastrového obrázku
	const width = 320
	const height = 240

	// jméno souboru s výsledným obrázkem
	const fileName = "03.png"

	// vytvoření kontextu
	dc := gg.NewContext(width, height)

	// vykreslování s využitím kontextu
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Clear()

	// uložení výsledného rastrového obrázku
	dc.SavePNG(fileName)
}
