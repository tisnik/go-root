// Seriál "Programovací jazyk Go"
//
// Patnáctá část
//
// Demonstrační příklad číslo 14:
//     Základní způsob použití knihovny GG

package main

import "github.com/fogleman/gg"

func main() {
	const width = 320
	const height = 240

	dc := gg.NewContext(width, height)

	dc.SetRGB(0.2, 0.0, 0.8)
	dc.DrawCircle(width/2, height/2, 100)

	dc.Fill()

	dc.SavePNG("14.png")
}
