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
// Demonstrační příklad číslo 22:
//    export do SVG, výsledek vykreslení gradientního lineárního přechodu
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_B0/22_linear_gradient_svg.html

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
	const fileName = "22.svg"

	// vykreslování s využitím kontextu
	color := gg.RGBA{R: 0.9, G: 0.9, B: 0.7, A: 1.0}
	recorder.ClearWithColor(color)

	// barvy použité pro výplně gradientním přechodem
	customColor1 := gg.Hex("44aadd")
	customColor2 := gg.Hex("dd44aa")
	customColor3 := gg.Hex("aadd44aa")

	const border = 40

	// lineární přechod mezi dvěma barvami
	gradientFill := gg.NewLinearGradientBrush(
		border, height/2, width-border*2, height/2).
		AddColorStop(0, customColor1).
		AddColorStop(0.7, customColor2).
		AddColorStop(1, customColor3)

	// použije se výše definoaný lineární přechod mezi dvěma barvami
	recorder.SetFillBrush(gradientFill)

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
