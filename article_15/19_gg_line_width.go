// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Patnáctá část
//     Programovací jazyk Go a grafika: tvorba animovaných GIFů, grafická knihovna GG
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-grafika-tvorba-animovanych-gifu-graficka-knihovna-gg/
//
// Demonstrační příklad číslo 19:
//     Specifikace šířky vykreslovaných cest
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_15/19_gg_line_width.html

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

	for i := 0; i < 256; i += 16 {
		width := float64(i) / 20
		dc.SetLineWidth(width)

		x := float64(i + 32)
		dc.DrawLine(x, 20, x, height-20)

		dc.Stroke()
	}

	dc.SavePNG("19.png")
}
