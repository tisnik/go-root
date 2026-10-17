// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sto čtrnáctá část
//    Tvorba 2D i 3D grafiky a animací v Go s využitím projektu GoGPU (2. část)
//    https://www.root.cz/clanky/tvorba-2d-i-3d-grafiky-a-animaci-v-go-s-vyuzitim-projektu-gogpu-2-cast/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze sto čtrnácté části:
//    https://github.com/tisnik/go-root/blob/master/article_B0/README.md
//
// Demonstrační příklad číslo 18:
//    gradient závislý na úhlu, přechod dvou barev
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_B0/18_sweep_gradient.html

package main

import "github.com/gogpu/gg"

func main() {
	// rozměry rastrového obrázku
	const width = 400
	const height = 300

	// jméno souboru s výsledným obrázkem
	const fileName = "18.png"

	// vytvoření kontextu
	dc := gg.NewContext(width, height)

	// vykreslování s využitím kontextu
	color := gg.RGBA{R: 0.9, G: 0.9, B: 0.7, A: 1.0}
	dc.ClearWithColor(color)

	// barvy použité pro výplně gradientním přechodem
	customColor1 := gg.Hex("dd4444")
	customColor2 := gg.Hex("ffffff")

	gradientFill := gg.NewSweepGradientBrush(width/2, height/2, 0).
		AddColorStop(0.0, customColor1).
		AddColorStop(1.0, customColor2)

	dc.SetFillBrush(gradientFill)

	const border = 0
	dc.DrawRectangle(border, border, width-border*2, height-border*2)
	dc.Fill()

	// uložení výsledného rastrového obrázku
	dc.SavePNG(fileName)
}
