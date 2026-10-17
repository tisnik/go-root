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
// Demonstrační příklad číslo 13:
//    lineární gradient, přechod dvou barev, odlišný úhel
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_B0/13_linear_gradient.html

package main

import "github.com/gogpu/gg"

func main() {
	// rozměry rastrového obrázku
	const width = 400
	const height = 300

	// jméno souboru s výsledným obrázkem
	const fileName = "13.png"

	// vytvoření kontextu
	dc := gg.NewContext(width, height)

	// vykreslování s využitím kontextu
	color := gg.RGBA{R: 0.9, G: 0.9, B: 0.7, A: 1.0}
	dc.ClearWithColor(color)

	// barvy použité pro výplně gradientním přechodem
	customColor1 := gg.Hex("44aadd")
	customColor2 := gg.Hex("dd44aa")

	const border = 40

	// lineární přechod mezi dvěma barvami
	gradientFill := gg.NewLinearGradientBrush(
		border, border, width-border*2, height-border*2).
		AddColorStop(0, customColor1).
		AddColorStop(1, customColor2)

	// použije se výše definoaný lineární přechod mezi dvěma barvami
	dc.SetFillBrush(gradientFill)

	dc.DrawRectangle(border, border, width-border*2, height-border*2)
	dc.Fill()

	// uložení výsledného rastrového obrázku
	dc.SavePNG(fileName)
}
