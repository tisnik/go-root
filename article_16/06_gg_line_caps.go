// Seriál "Programovací jazyk Go"
//
// Šestnáctá část
//     Programovací jazyk Go a grafika: další užitečné funkce poskytované knihovnou GG
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-grafika-dalsi-uzitecne-funkce-poskytovane-knihovnou-gg/
//
// Demonstrační příklad číslo 6:
//     Nastavení stylu konců úseček.

package main

import "github.com/fogleman/gg"

func main() {
	const width = 320
	const height = 240

	dc := gg.NewContext(width, height)

	dc.DrawRectangle(0, 0, width, height)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()

	dc.SetRGBA(0.0, 0.0, 0.0, 1.0)

	dc.SetLineWidth(10.0)

	for i := 0; i < 256; i += 16 {
		x := float64(i + 32)

		width := float64(i) / 20
		dc.SetLineWidth(width)

		dc.SetLineCapRound()
		dc.DrawLine(x, 20, x, 75)
		dc.Stroke()

		dc.SetLineCapButt()
		dc.DrawLine(x, 92, x, height-92)
		dc.Stroke()

		dc.SetLineCapSquare()
		dc.DrawLine(x, height-75, x, height-20)
		dc.Stroke()
	}

	dc.SetRGBA(1.0, 0.0, 0.0, 1.0)
	dc.SetLineWidth(1.0)

	dc.DrawLine(32, 20, width-32, 20)
	dc.Stroke()

	dc.DrawLine(32, 75, width-32, 75)
	dc.Stroke()

	dc.DrawLine(32, 92, width-32, 92)
	dc.Stroke()

	dc.DrawLine(32, height-92, width-32, height-92)
	dc.Stroke()

	dc.DrawLine(32, height-75, width-32, height-75)
	dc.Stroke()

	dc.DrawLine(32, height-20, width-32, height-20)
	dc.Stroke()

	dc.SavePNG("06.png")
}
