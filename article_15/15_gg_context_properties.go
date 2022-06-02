// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Patnáctá část
//     Programovací jazyk Go a grafika: tvorba animovaných GIFů, grafická knihovna GG
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-grafika-tvorba-animovanych-gifu-graficka-knihovna-gg/
//
// Seznam příkladů z patnácté části:
//    https://github.com/tisnik/go-root/blob/master/article_15/README.md
//
// Demonstrační příklad číslo 15:
//     Použití kontextu a nastavení vlastností kreslení
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_15/15_gg_context_properties.html

package main

import "github.com/fogleman/gg"

func main() {
	const width = 320
	const height = 240

	dc := gg.NewContext(width, height)

	dc.DrawCircle(width/2, height/2, 100)

	dc.SetRGB(0.2, 0.0, 0.8)
	dc.Fill()

	dc.SavePNG("15.png")
}
