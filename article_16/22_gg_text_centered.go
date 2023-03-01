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
// Demonstrační příklad číslo 22:
//     Vykreslení vycentrovaného textu.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_16/22_gg_text_centered.html

package main

import "github.com/fogleman/gg"

func main() {
	const width = 320
	const height = 240

	dc := gg.NewContext(width, height)

	dc.DrawRectangle(0, 0, width, height)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()

	dc.SetRGB(0.0, 0.0, 0.0)
	if err := dc.LoadFontFace("luxisr.ttf", 36); err != nil {
		println("Cannot load font")
		panic(err)
	}
	dc.DrawStringAnchored("Hello, world!", width/2, height/2, 0.5, 0.5)

	dc.SavePNG("22.png")
}
