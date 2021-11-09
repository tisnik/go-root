// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Patnáctá část
//     Programovací jazyk Go a grafika: tvorba animovaných GIFů, grafická knihovna GG
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-grafika-tvorba-animovanych-gifu-graficka-knihovna-gg/
//
// Demonstrační příklad číslo 16:
//     Vykreslení cesty
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_15/16_gg_path_stroke.html

package main

import "github.com/fogleman/gg"

func main() {
	const width = 320
	const height = 240

	dc := gg.NewContext(width, height)

	dc.SetRGB(0.2, 0.0, 0.8)
	dc.DrawLine(10, 10, width-10, height-10)
	dc.Stroke()

	dc.SetRGB(0.8, 0.0, 0.2)
	dc.DrawLine(10, height-10, width-10, 10)
	dc.Stroke()

	dc.SavePNG("16.png")
}
