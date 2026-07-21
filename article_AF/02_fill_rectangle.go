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
// Seznam demonstračních příkladů ze čtrnácté části:
//    https://github.com/tisnik/go-root/blob/master/article_14/README.md
//
// Demonstrační příklad číslo 2:
//    Balíček gogpu/gg.
//    Vymazání obrázku vykreslením vyplněného obdélníku.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_A0/02_fill_rectangle.html

package main

import "github.com/gogpu/gg"

func main() {
	// rozměry rastrového obrázku
	const width = 320
	const height = 240

	// jméno souboru s výsledným obrázkem
	const fileName = "02.png"

	// vytvoření kontextu
	dc := gg.NewContext(width, height)

	// vykreslování s využitím kontextu
	dc.DrawRectangle(0, 0, width, height)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()

	// uložení výsledného rastrového obrázku
	dc.SavePNG(fileName)
}
