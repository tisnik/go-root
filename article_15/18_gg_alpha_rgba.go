// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Patnáctá část
//     Programovací jazyk Go a grafika: tvorba animovaných GIFů, grafická knihovna GG
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-grafika-tvorba-animovanych-gifu-graficka-knihovna-gg/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z patnácté části:
//    https://github.com/tisnik/go-root/blob/master/article_15/README.md
//
// Demonstrační příklad číslo 18:
//     Specifikace alfa kanálu (průhlednosti)
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_15/18_gg_alpha_rgba.html

package main

import "github.com/fogleman/gg"

func main() {
	const width = 320
	const height = 240

	dc := gg.NewContext(width, height)

	dc.DrawRectangle(0, 0, width, height)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()

	for i := 0; i < 256; i += 8 {
		// průhlednost je hodnota v rozsahu 0.0 až 1.0
		alpha := float64(i) / 256.0
		dc.SetRGBA(0.0, 0.0, 0.0, alpha)

		x := float64(i + 32)
		dc.DrawLine(x, 20, x, height-20)

		dc.Stroke()
	}

	dc.SavePNG("18.png")
}
