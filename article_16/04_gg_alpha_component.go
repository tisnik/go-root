// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestnáctá část
//     Programovací jazyk Go a grafika: další užitečné funkce poskytované knihovnou GG
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-grafika-dalsi-uzitecne-funkce-poskytovane-knihovnou-gg/
//
// Seznam příkladů ze šestnácté části:
//    https://github.com/tisnik/go-root/blob/master/article_16/README.md
//
// Demonstrační příklad číslo 4:
//     Alfa kanál v barvovém prostoru RGBA.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_16/04_gg_alpha_component.html

package main

import "github.com/fogleman/gg"

func main() {
	const width = 320
	const height = 240

	dc := gg.NewContext(width, height)

	dc.DrawRectangle(0, 0, width, height/2)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()

	dc.DrawRectangle(0, height/2, width, height)
	dc.SetRGB(0.0, 0.0, 0.0)
	dc.Fill()

	for i := 0; i < 256; i += 8 {
		x := float64(i + 32)

		// průhlednost je hodnota v rozsahu 0.0 až 1.0
		alpha := float64(i) / 256.0
		dc.SetRGBA(0.0, 0.0, 0.0, alpha)

		dc.DrawLine(x, 20, x, height/2-20)

		dc.Stroke()

		dc.SetRGBA(1.0, 1.0, 1.0, alpha)
		dc.DrawLine(x, height/2+20, x, height-20)

		dc.Stroke()
	}

	dc.SavePNG("04.png")
}
