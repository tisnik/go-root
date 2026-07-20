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
	const fileName = "20.svg"

	// vykreslování s využitím kontextu
	color := gg.RGBA{R: 0.9, G: 0.9, B: 0.7, A: 1.0}
	recorder.ClearWithColor(color)

	const border = 40
	recorder.DrawRectangle(border, border, width-border*2, height-border*2)
	recorder.SetRGB(0.0, 0.0, 0.0)

	// vykreslení obrysu
	recorder.Stroke()

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
