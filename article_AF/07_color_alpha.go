// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sto třináctá část
//    Tvorba 2D i 3D grafiky a animací v Go s využitím projektu GoGPU
//    https://www.root.cz/clanky/tvorba-2d-i-3d-grafiky-a-animaci-v-go-s-vyuzitim-projektu-gogpu/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze sto třinácté části:
//    https://github.com/tisnik/go-root/blob/master/article_AF/README.md
//
// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Demonstrační příklad číslo 7:
//    Balíček gogpu/gg.
//    Práce s barvami.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_AF/07_color_alpha.html

package main

import "github.com/gogpu/gg"

func main() {
	// rozměry rastrového obrázku
	const width = 256
	const height = 256

	// jméno souboru s výsledným obrázkem
	const fileName = "07.png"

	// vytvoření kontextu
	dc := gg.NewContext(width, height)

	for y := range 16 {
		for x := range 16 {
			// nastavení barvy kreslení
			dc.SetRGBA(0.0, 0.0, 0.0, float64(x)/32.0+float64(y)/32.0)
			// vykreslení vyplněného obdélníka
			dc.DrawRectangle(float64(x*16)+3, float64(y*16)+3, 10, 10)
			dc.Fill()
		}
	}

	// uložení výsledného rastrového obrázku
	dc.SavePNG(fileName)
}
