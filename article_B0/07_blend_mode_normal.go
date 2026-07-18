package main

import "github.com/gogpu/gg"

func main() {
	// rozměry rastrového obrázku
	const width = 400
	const height = 300

	// jméno souboru s výsledným obrázkem
	const fileName = "07.png"

	// vytvoření kontextu
	dc := gg.NewContext(width, height)

	// vykreslování s využitím kontextu
	color := gg.RGBA{R: 0.9, G: 0.9, B: 0.7, A: 1.0}
	dc.ClearWithColor(color)

	const border = 40

	customColor1 := gg.Hex("44aadd")
	dc.SetFillBrush(gg.Solid(customColor1))

	dc.DrawRectangle(border, border, width-border*4, height-border*4)

	// vyplnění obrysu
	dc.Fill()

	customColor2 := gg.Hex("ddaa44")
	dc.SetFillBrush(gg.Solid(customColor2))

	dc.DrawRectangle(border*2, border*2, width-border*4, height-border*4)

	// vyplnění obrysu
	dc.Fill()

	customColor3 := gg.Hex("dd44aa")
	dc.SetFillBrush(gg.Solid(customColor3))

	dc.DrawRectangle(border*3, border*3, width-border*4, height-border*4)

	// vyplnění obrysu
	dc.Fill()

	// uložení výsledného rastrového obrázku
	dc.SavePNG(fileName)
}
