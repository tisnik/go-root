package main

import (
	"math"

	"github.com/gogpu/gg"
)

func main() {
	// rozměry rastrového obrázku
	const width = 400
	const height = 300

	// jméno souboru s výsledným obrázkem
	const fileName = "10.png"

	// vytvoření kontextu
	dc := gg.NewContext(width, height)

	// vykreslování s využitím kontextu
	color := gg.RGBA{R: 0.9, G: 0.9, B: 0.7, A: 1.0}
	dc.ClearWithColor(color)

	customColor := gg.Hex("44aadd")
	dc.SetFillBrush(gg.Solid(customColor))

	const radius = height / 3

	// začátek vykreslování cesty
	dc.MoveTo(width/2, height/2-radius)
	for i := range 4 {
		angle := float64(i+1) * 3.0 * 360 / 5.0
		x := radius * math.Sin(angle*math.Pi/180.0)
		y := radius * math.Cos(angle*math.Pi/180.0)
		dc.LineTo(width/2+x, height/2-y)
	}
	dc.ClosePath()

	// vyplnění obrysu
	dc.Fill()

	// uložení výsledného rastrového obrázku
	dc.SavePNG(fileName)
}
