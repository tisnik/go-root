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
// Demonstrační příklad číslo 10:
//    výchozí pravidlo pro vyplňování obrazců, jejichž hrany se překrývají
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_B0/10_fill_rule_default.html

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
