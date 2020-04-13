// Seriál "Programovací jazyk Go"
//
// Šestnáctá část
//     Programovací jazyk Go a grafika: další užitečné funkce poskytované knihovnou GG
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-grafika-dalsi-uzitecne-funkce-poskytovane-knihovnou-gg/
//
// Demonstrační příklad číslo 8:
//     Práce s cestou (path).

package main

import "github.com/fogleman/gg"

func main() {
	const width = 320
	const height = 240

	dc := gg.NewContext(width, height)

	dc.DrawRectangle(0, 0, width, height)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()

	dc.SetRGBA(0.0, 0.0, 0.0, 1)

	dc.MoveTo(100, 200)
	dc.LineTo(200, 200)
	dc.LineTo(100, 100)
	dc.LineTo(100, 200)
	dc.LineTo(200, 100)
	dc.LineTo(100, 100)
	dc.LineTo(150, 50)
	dc.LineTo(200, 100)
	dc.LineTo(200, 200)

	dc.Stroke()

	dc.SavePNG("08.png")
}
