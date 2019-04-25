// Seriál "Programovací jazyk Go"
//
// Patnáctá část
//
// Demonstrační příklad číslo 15:
//     Použití kontextu a nastavení vlastností kreslení

package main

import "github.com/fogleman/gg"

func main() {
	const width = 320
	const height = 240

	dc := gg.NewContext(width, height)

	dc.DrawCircle(width/2, height/2, 100)

	dc.SetRGB(0.2, 0.0, 0.8)
	dc.Fill()

	dc.SavePNG("15.png")
}
