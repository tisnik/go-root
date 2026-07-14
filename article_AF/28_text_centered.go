// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Demonstrační příklad číslo 28:
//    Balíček gogpu/gg.
//    Vykreslení vycentrovaného textu.
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
	const fileName = "28.png"

	// jméno souboru s fontem
	const fontFileName = "fonts/FreeSans.ttf"

	// vytvoření kontextu
	dc := gg.NewContext(width, height)

	dc.DrawRectangle(0, 0, width, height)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()

	dc.SetRGB(0.0, 0.0, 0.0)

	// pokus o načtení fontu
	if err := dc.LoadFontFace(fontFileName, 24); err != nil {
		fmt.Println("Cannot load font")
		panic(err)
	}

	var h float64 = 20
	weight := -1.0

	for i := 0; i < 11; i++ {
		dc.DrawStringAnchored("Hello, world!", width/2, h, weight, 0.5)
		h += 20
		weight += 0.2
	}

	dc.SavePNG(fileName)
}
