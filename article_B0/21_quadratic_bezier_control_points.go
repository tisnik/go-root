package main

import (
	"log"

	"github.com/gogpu/gg"
	_ "github.com/gogpu/gg-svg"
	"github.com/gogpu/gg/recording"
)

func drawQuadraticBezier(dc *recording.Recorder, x0 float64, y0 float64, x1 float64, y1 float64, x2 float64, y2 float64) {
	dc.SetRGBA(1.0, 0.5, 0.5, 1.0)
	dc.DrawLine(x0, y0, x1, y1)
	dc.Stroke()
	dc.DrawLine(x1, y1, x2, y2)
	dc.Stroke()

	dc.SetRGBA(0.0, 0.0, 0.0, 1.0)
	dc.MoveTo(x0, y0)
	dc.QuadraticTo(x1, y1, x2, y2)
	dc.Stroke()

	dc.SetRGBA(0.2, 0.2, 1.0, 1.0)

	dc.DrawCircle(x0, y0, 3)
	// vykreslení
	dc.Stroke()

	dc.DrawCircle(x1, y1, 3)
	// vykreslení
	dc.Stroke()

	dc.DrawCircle(x2, y2, 3)
	// vykreslení
	dc.Stroke()
}

func main() {
	// rozměry obrázku
	const width = 400
	const height = 300

	// objekt realizující zápis zaznamenaných kreslicích příkazů do SVG
	recorder := recording.NewRecorder(width, height)

	// jméno souboru s výslednou kresbou
	const fileName = "21.svg"

	// vykreslování s využitím kontextu
	color := gg.RGBA{R: 0.9, G: 0.9, B: 0.7, A: 1.0}
	recorder.ClearWithColor(color)

	// vykreslení cest
	drawQuadraticBezier(recorder, 10, 150, 50, 10, 90, 150)
	drawQuadraticBezier(recorder, 110, 100, 190, 100, 150, 190)
	drawQuadraticBezier(recorder, 250, 150, 210, 60, 290, 150)

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
