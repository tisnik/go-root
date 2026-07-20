// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Demonstrační příklad číslo 26:
//    Balíček gogpu/gg.
//    Vykreslení textu.
//

package main

import (
	"fmt"

	"github.com/gogpu/gg"
)

func main() {
	// rozměry rastrového obrázku
	const width = 320
	const height = 240

	// jméno souboru s výsledným obrázkem
	const fileName = "26.png"

	// jméno souboru s fontem
	const fontFileName = "fonts/FreeSans.ttf"

	// vytvoření kontextu
	dc := gg.NewContext(width, height)

	dc.DrawRectangle(0, 0, width, height)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()

	dc.SetRGB(0.0, 0.0, 0.0)

	// pokus o načtení fontu
	if err := dc.LoadFontFace(fontFileName, 36); err != nil {
		fmt.Println("Cannot load font")
		panic(err)
	}
	dc.DrawString("Hello, world!", 0, height)

	// uložení výsledného rastrového obrázku
	dc.SavePNG(fileName)
}
