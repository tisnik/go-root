package main

import (
	"fmt"

	"github.com/gogpu/gg"
	_ "github.com/gogpu/gg/gpu"
	"github.com/gogpu/gg/integration/ggcanvas"
	"github.com/gogpu/gogpu"
)

func main() {
	// rozměry okna
	const width = 400
	const height = 300

	// inicializace aplikace s GUI
	app := gogpu.NewApp(gogpu.DefaultConfig().
		WithTitle("GoGPU + gg").
		WithSize(width, height))

	var canvas *ggcanvas.Canvas

	// callback metoda zavolaná ve chvíli, kdy se má překreslit okno
	app.OnDraw(func(dc *gogpu.Context) {
		fmt.Println("Redraw....")
		if canvas == nil {
			provider := app.GPUContextProvider()
			w, h := dc.Width(), dc.Height()
			canvas, _ = ggcanvas.New(provider, w, h)
		}

		cw, ch := canvas.Size()

		canvas.Draw(func(cc *gg.Context) {
			renderFrame(cc, cw, ch)
		})

		canvas.Render(dc.RenderTarget())
	})

	// callback metoda zavolaná ve chvíli, kdy se okno uzavírá
	app.OnClose(func() {
		fmt.Println("Closing")
		gg.CloseAccelerator()
	})

	// spuštění aplikace
	app.Run()
}

func renderFrame(dc *gg.Context, width, height int) {
	// barva pozadí
	dc.SetRGB(0.9, 0.9, 0.7)

	// obdélník s plochou odpovídající oknu
	dc.DrawRectangle(0, 0, float64(width), float64(height))
	dc.Fill()

	const border = 40
	dc.DrawRectangle(border, border, float64(width)-border*2, float64(height)-border*2)
	dc.SetRGB(0.0, 0.0, 0.0)

	// vykreslení obrysu
	dc.Stroke()
}
