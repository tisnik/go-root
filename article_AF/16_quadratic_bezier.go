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
// Demonstrační příklad číslo 16:
//    Balíček gogpu/gg.
//    Kvadratická Bézierova křivka.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_AF/16_quadratic_bezier.html

package main

import "github.com/gogpu/gg"

func main() {
	// rozměry rastrového obrázku
	const width = 320
	const height = 240

	// jméno souboru s výsledným obrázkem
	const fileName = "16.png"

	// vytvoření kontextu
	dc := gg.NewContext(width, height)

	// pozadí obrázku
	dc.DrawRectangle(0, 0, width, height)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()

	// cesta
	dc.SetRGBA(1.0, 0.0, 0.0, 1.0)
	dc.MoveTo(10, 150)
	dc.QuadraticTo(50, 10, 90, 150)
	// vlastní vykreslení cesty
	dc.Stroke()

	// cesta
	dc.SetRGBA(0.0, 1.0, 0.0, 1.0)
	dc.MoveTo(110, 100)
	dc.QuadraticTo(190, 100, 150, 190)
	// vlastní vykreslení cesty
	dc.Stroke()

	// cesta
	dc.SetRGBA(0.0, 0.0, 1.0, 1.0)
	dc.MoveTo(250, 150)
	dc.QuadraticTo(210, 60, 290, 150)
	// vlastní vykreslení cesty
	dc.Stroke()

	// uložení výsledného rastrového obrázku
	dc.SavePNG(fileName)
}
