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
// Demonstrační příklad číslo 9:
//    Balíček gogpu/gg.
//    Barvový prostor RGB a RGBA.
//    Alfa kanál v barvovém prostoru RGBA.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_AF/09_alpha_component.html

package main

import "github.com/gogpu/gg"

func main() {
	// rozměry rastrového obrázku
	const width = 320
	const height = 240

	// jméno souboru s výsledným obrázkem
	const fileName = "09.png"

	// vytvoření kontextu
	dc := gg.NewContext(width, height)

	// pozadí obrázku
	dc.DrawRectangle(0, 0, width, height/2)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()

	// pozadí obrázku
	dc.DrawRectangle(0, height/2, width, height)
	dc.SetRGB(0.0, 0.0, 0.0)
	dc.Fill()

	for i := 0; i < 256; i += 8 {
		x := float64(i + 32)

		// průhlednost je hodnota v rozsahu 0.0 až 1.0
		alpha := float64(i) / 256.0
		dc.SetRGBA(0.0, 0.0, 0.0, alpha)

		dc.DrawLine(x, 20, x, height/2-20)

		dc.Stroke()

		dc.SetRGBA(1.0, 1.0, 1.0, alpha)
		dc.DrawLine(x, height/2+20, x, height-20)
		dc.Stroke()
	}

	// uložení výsledného rastrového obrázku
	dc.SavePNG(fileName)
}
