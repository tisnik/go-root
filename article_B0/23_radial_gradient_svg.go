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
// Demonstrační příklad číslo 23:
//    export do SVG, výsledek vykreslení gradientního radiálního přechodu
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_B0/23_radial_gradient_svg.html

package main

import (
	"log"

	"github.com/gogpu/gg"
	_ "github.com/gogpu/gg-svg"
	"github.com/gogpu/gg/recording"
)

func main() {
	// rozměry obrázku
	const width = 400
	const height = 300

	// objekt realizující zápis zaznamenaných kreslicích příkazů do SVG
	recorder := recording.NewRecorder(width, height)

	// jméno souboru s výslednou kresbou
	const fileName = "23.svg"

	// vykreslování s využitím kontextu
	color := gg.RGBA{R: 0.9, G: 0.9, B: 0.7, A: 1.0}
	recorder.ClearWithColor(color)

	// barvy použité pro výplně gradientním přechodem
	customColor1 := gg.Hex("dd4444")
	customColor2 := gg.Hex("000000")
	customColor3 := gg.Hex("ffffff")

	gradientFill := gg.NewRadialGradientBrush(width/2, height/2, 0, height/2).
		AddColorStop(0.0, customColor1).
		AddColorStop(0.4, customColor2).
		AddColorStop(1.0, customColor3)

	recorder.SetFillBrush(gradientFill)

	// použije se výše definoaný lineární přechod mezi dvěma barvami
	recorder.SetFillBrush(gradientFill)

	const border = 0
	recorder.DrawRectangle(border, border, width-border*2, height-border*2)
	recorder.Fill()

	// dokončení záznamu
	drawing_recording := recorder.FinishRecording()

	// backed pro zápis do SVG
	backend, err := recording.NewBackend("svg")
	if err != nil {
		log.Fatal(err)
	}

	// provedení zápisu do SVG
	drawing_recording.Playback(backend)

	// uložení výsledného vektorového obrázku
	if fb, ok := backend.(recording.FileBackend); ok {
		fb.SaveToFile(fileName)
	}

}
