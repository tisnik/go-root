// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestnáctá část
//     Programovací jazyk Go a grafika: další užitečné funkce poskytované knihovnou GG
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-grafika-dalsi-uzitecne-funkce-poskytovane-knihovnou-gg/
//
// Demonstrační příklad číslo 23:
//     Vykreslení vycentrovaného textu.

package main

import "github.com/fogleman/gg"

func main() {
	const width = 320
	const height = 240

	dc := gg.NewContext(width, height)

	dc.DrawRectangle(0, 0, width, height)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()

	dc.SetRGB(0.0, 0.0, 0.0)
	if err := dc.LoadFontFace("luxisr.ttf", 24); err != nil {
		println("Cannot load font")
		panic(err)
	}

	var h float64 = 20
	weight := -1.0

	for i := 0; i < 11; i++ {
		dc.DrawStringAnchored("Hello, world!", width/2, h, weight, 0.5)
		h += 20
		weight += 0.2
	}

	dc.SavePNG("23.png")
}
