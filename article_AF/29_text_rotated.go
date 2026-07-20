// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Demonstrační příklad číslo 29:
//    Balíček gogpu/gg.
//    Vykreslení vycentrovaného textu.
//

package main

import (
	"fmt"
	"math"

	"github.com/gogpu/gg"
)

func Radians(angle float64) float64 {
	return float64(angle / 180.0 * math.Pi)
}

func main() {
	// rozměry rastrového obrázku
	const width = 320
	const height = 240

	// jméno souboru s výsledným obrázkem
	const fileName = "29.png"

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
	// transformace a vykreslení
	dc.Push()
	dc.RotateAbout(Radians(0.0), width/2, height/2)
	dc.DrawStringAnchored("Hello, world!", width/2, height/2, 0.5, 0.5)
	dc.RotateAbout(Radians(120.0), width/2, height/2)
	dc.DrawStringAnchored("Hello, world!", width/2, height/2, 0.5, 0.5)
	dc.RotateAbout(Radians(120.0), width/2, height/2)
	dc.DrawStringAnchored("Hello, world!", width/2, height/2, 0.5, 0.5)
	dc.Pop()

	// uložení výsledného rastrového obrázku
	dc.SavePNG(fileName)
}
