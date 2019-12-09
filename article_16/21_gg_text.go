// Seriál "Programovací jazyk Go"
//
// Šestnáctá část
//
// Demonstrační příklad číslo 21:
//     Vykreslení textu.

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
	if err := dc.LoadFontFace("luxisr.ttf", 36); err != nil {
		println("Cannot load font")
		panic(err)
	}
	dc.DrawString("Hello, world!", 0, height)

	dc.SavePNG("21.png")
}
