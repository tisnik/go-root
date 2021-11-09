// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestnáctá část
//     Programovací jazyk Go a grafika: další užitečné funkce poskytované knihovnou GG
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-grafika-dalsi-uzitecne-funkce-poskytovane-knihovnou-gg/
//
// Demonstrační příklad číslo 3:
//     Barvový prostor RGB a RGBA.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_16/03_gg_rgba_color.html

package main

import "github.com/fogleman/gg"

func main() {
	const width = 320
	const height = 240

	dc := gg.NewContext(width, height)

	dc.DrawRectangle(0, 0, width, height)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()

	for i := 0; i < 256; i += 4 {
		x := float64(i + 32)

		// barvové složky jsou hodnoty v rozsahu 0.0 až 1.0
		r := float64(i) / 256.0
		dc.SetRGBA(r, 0.0, 0.0, 1.0)
		dc.DrawLine(x, 20, x, 75)
		dc.Stroke()

		g := float64(i) / 256.0
		dc.SetRGBA(0.0, g, 0.0, 1.0)
		dc.DrawLine(x, 92, x, height-92)
		dc.Stroke()

		b := float64(i) / 256.0
		dc.SetRGBA(0.0, 0.0, b, 1.0)
		dc.DrawLine(x, height-75, x, height-20)
		dc.Stroke()
	}

	dc.SavePNG("03.png")
}
