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
// Demonstrační příklad číslo 10:
//     Uzavřená cesta (closed path).
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_16/10_gg_closed_path.html

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
	dc.LineTo(100, 100)
	dc.LineTo(150, 50)
	dc.LineTo(200, 100)
	dc.LineTo(200, 200)
	dc.ClosePath()

	dc.Stroke()

	dc.SavePNG("10.png")
}
