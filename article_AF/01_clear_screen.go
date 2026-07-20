// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Demonstrační příklad číslo 1:
//    Balíček gogpu/gg.
//    Vymazání obsahu rastrového obrázku.
//

package main

import "github.com/gogpu/gg"

func main() {
	// rozměry rastrového obrázku
	const width = 320
	const height = 240

	// jméno souboru s výsledným obrázkem
	const fileName = "01.png"

	// vytvoření kontextu
	dc := gg.NewContext(width, height)

	// smazání obsahu obrázku s využitím kontextu
	dc.Clear()

	// uložení výsledného rastrového obrázku
	dc.SavePNG(fileName)
}
