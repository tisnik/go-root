// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestnáctá část
//     Programovací jazyk Go a grafika: další užitečné funkce poskytované knihovnou GG
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-grafika-dalsi-uzitecne-funkce-poskytovane-knihovnou-gg/
//
// Demonstrační příklad číslo 2:
//     Vymazání obrazovky (druhá, lepší varianta).
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_16/02_gg_clear_screen.html

package main

import "github.com/fogleman/gg"

func main() {
	const width = 320
	const height = 240

	dc := gg.NewContext(width, height)

	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Clear()

	dc.SavePNG("02.png")
}
