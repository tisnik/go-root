// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Patnáctá část
//     Programovací jazyk Go a grafika: tvorba animovaných GIFů, grafická knihovna GG
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-grafika-tvorba-animovanych-gifu-graficka-knihovna-gg/
//
// Demonstrační příklad číslo 14:
//     Základní způsob použití knihovny GG
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_15/14_gg_basic.html

package main

import "github.com/fogleman/gg"

func main() {
	const width = 320
	const height = 240

	dc := gg.NewContext(width, height)

	dc.SetRGB(0.2, 0.0, 0.8)
	dc.DrawCircle(width/2, height/2, 100)

	dc.Fill()

	dc.SavePNG("14.png")
}
