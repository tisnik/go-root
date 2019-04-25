// Seriál "Programovací jazyk Go"
//
// Patnáctá část
//
// Demonstrační příklad číslo 18:
//     Specifikace alfa kanálu (průhlednosti)

package main

import "github.com/fogleman/gg"

func main() {
	const width = 320
	const height = 240

	dc := gg.NewContext(width, height)

	dc.DrawRectangle(0, 0, width, height)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()

	for i := 0; i < 256; i += 8 {
		// průhlednost je hodnota v rozsahu 0.0 až 1.0
		alpha := float64(i) / 256.0
		dc.SetRGBA(0.0, 0.0, 0.0, alpha)

		x := float64(i + 32)
		dc.DrawLine(x, 20, x, height-20)

		dc.Stroke()
	}

	dc.SavePNG("18.png")
}
