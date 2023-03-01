// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestnáctá část
//     Programovací jazyk Go a grafika: další užitečné funkce poskytované knihovnou GG
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-grafika-dalsi-uzitecne-funkce-poskytovane-knihovnou-gg/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze šestnácté části:
//    https://github.com/tisnik/go-root/blob/master/article_16/README.md
//
// Demonstrační příklad číslo 11:
//     Kvadratická Bézierova křivka.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_16/11_gg_quadratic_bezier.html

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
