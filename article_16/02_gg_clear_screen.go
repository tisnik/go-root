// Seriál "Programovací jazyk Go"
//
// Šestnáctá část
//     Programovací jazyk Go a grafika: další užitečné funkce poskytované knihovnou GG
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-grafika-dalsi-uzitecne-funkce-poskytovane-knihovnou-gg/
//
// Demonstrační příklad číslo 2:
//     Vymazání obrazovky (druhá, lepší varianta).

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
