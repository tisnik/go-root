// Seriál "Programovací jazyk Go"
//
// Šestnáctá část
//     Programovací jazyk Go a grafika: další užitečné funkce poskytované knihovnou GG
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-grafika-dalsi-uzitecne-funkce-poskytovane-knihovnou-gg/
//
// Demonstrační příklad číslo 11:
//     Kvadratická Bézierova křivka.

package main

import "github.com/fogleman/gg"

func main() {
	const width = 320
	const height = 240

	dc := gg.NewContext(width, height)

	dc.DrawRectangle(0, 0, width, height)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()

	dc.SetRGBA(1.0, 0.0, 0.0, 1.0)
	dc.MoveTo(10, 150)
	dc.QuadraticTo(50, 10, 90, 150)
	dc.Stroke()

	dc.SetRGBA(0.0, 1.0, 0.0, 1.0)
	dc.MoveTo(110, 100)
	dc.QuadraticTo(190, 100, 150, 190)
	dc.Stroke()

	dc.SetRGBA(0.0, 0.0, 1.0, 1.0)
	dc.MoveTo(250, 150)
	dc.QuadraticTo(210, 60, 290, 150)
	dc.Stroke()

	dc.SavePNG("11.png")
}
