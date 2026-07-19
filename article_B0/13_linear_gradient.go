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
