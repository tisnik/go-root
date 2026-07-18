package main

import "github.com/gogpu/gg"

func main() {
	// rozměry rastrového obrázku
	const width = 400
	const height = 300

	// jméno souboru s výsledným obrázkem
	const fileName = "06.png"

	// vytvoření kontextu
	dc := gg.NewContext(width, height)

	// vykreslování s využitím kontextu
	color := gg.RGBA{R: 0.9, G: 0.9, B: 0.7, A: 1.0}
	dc.ClearWithColor(color)

	blackColor := gg.Hex("000000")
	dc.SetStrokeBrush(gg.Solid(blackColor))

	customColor := gg.Hex("44aadd")
	dc.SetFillBrush(gg.Solid(customColor))

	const border = 40
	dc.DrawRectangle(border, border, width-border*2, height-border*2)

	// vykreslení obrysu
	dc.Fill()
	dc.Stroke()

	// uložení výsledného rastrového obrázku
	dc.SavePNG(fileName)
}
