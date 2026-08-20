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
// Demonstrační příklad číslo 23:
//    Balíček gogpu/gg.
//    Transformace (rotace okolo zvoleného bodu).
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_AF/23_transformation_rotate_about.html

package main

import (
	"math"

	"github.com/gogpu/gg"
)

func Radians(angle float64) float64 {
	return float64(angle / 180.0 * math.Pi)
}

func main() {
	// rozměry rastrového obrázku
	const width = 320
	const height = 240

	// jméno souboru s výsledným obrázkem
	const fileName = "23.png"

	// vytvoření kontextu
	dc := gg.NewContext(width, height)

	dc.DrawRectangle(0, 0, width, height)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()

	dc.RotateAbout(Radians(15.0), width/2.0, height/2.0)

	dc.SetRGBA(0.0, 0.0, 0.0, 1)

	dc.DrawLine(32, 20, 288, 20)
	dc.Stroke()

	dc.SetDash(10)
	dc.DrawLine(32, 40, 288, 40)
	dc.Stroke()

	dc.SetDash(10, 10)
	dc.DrawLine(32, 60, 288, 60)
	dc.Stroke()

	dc.SetDash(10, 5)
	dc.DrawLine(32, 80, 288, 80)
	dc.Stroke()

	dc.SetDash(10, 5, 2, 5)
	dc.DrawLine(32, 100, 288, 100)
	dc.Stroke()

	dc.SetLineWidth(4.0)
	dc.SetLineCap(gg.LineCapButt)

	dc.DrawLine(32, 140, 288, 140)
	dc.Stroke()

	dc.SetDash(10)
	dc.DrawLine(32, 160, 288, 160)
	dc.Stroke()

	dc.SetDash(10, 10)
	dc.DrawLine(32, 180, 288, 180)
	dc.Stroke()

	dc.SetDash(10, 5)
	dc.DrawLine(32, 200, 288, 200)
	dc.Stroke()

	dc.SetDash(10, 5, 2, 5)
	dc.DrawLine(32, 220, 288, 220)
	dc.Stroke()

	// uložení výsledného rastrového obrázku
	dc.SavePNG(fileName)
}
