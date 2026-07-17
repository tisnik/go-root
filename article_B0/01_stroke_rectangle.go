package main

import "github.com/gogpu/gg"

func main() {
	// rozměry rastrového obrázku
	const width = 400
	const height = 300

	// jméno souboru s výsledným obrázkem
	const fileName = "01.png"

	// vytvoření kontextu
	dc := gg.NewContext(width, height)

	// vykreslování s využitím kontextu
	color := gg.RGBA{R: 0.9, G: 0.9, B: 0.7, A: 1.0}
	dc.ClearWithColor(color)

	const border = 40
	dc.DrawRectangle(border, border, width-border*2, height-border*2)
	dc.SetRGB(0.0, 0.0, 0.0)

	// vykreslení obrysu
	dc.Stroke()

	// uložení výsledného rastrového obrázku
	dc.SavePNG(fileName)
}
