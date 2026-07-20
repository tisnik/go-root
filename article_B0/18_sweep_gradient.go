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
