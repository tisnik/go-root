package main

import (
	"fmt"

	"github.com/gogpu/gg"
	_ "github.com/gogpu/gg/gpu"
	"github.com/gogpu/gg/integration/ggcanvas"
	"github.com/gogpu/gogpu"
)

func drawCubicBezier(dc *gg.Context,
	x0 float64, y0 float64, x1 float64, y1 float64,
	x2 float64, y2 float64, x3 float64, y3 float64) {

	dc.SetRGBA(1.0, 0.5, 0.5, 1.0)
	dc.DrawLine(x0, y0, x1, y1)
	dc.Stroke()
	dc.DrawLine(x1, y1, x2, y2)
	dc.Stroke()
	dc.DrawLine(x2, y2, x3, y3)
	dc.Stroke()

	dc.SetRGBA(0.0, 0.0, 0.0, 1.0)
	dc.MoveTo(x0, y0)
	dc.CubicTo(x1, y1, x2, y2, x3, y3)
	dc.Stroke()

	dc.SetRGBA(0.2, 0.2, 1.0, 1.0)

	dc.DrawCircle(x0, y0, 3)
	dc.Stroke()

	dc.DrawCircle(x1, y1, 3)
	dc.Stroke()

	dc.DrawCircle(x2, y2, 3)
	// vykreslení
	dc.Stroke()

	dc.DrawCircle(x2, y2, 3)
	// vykreslení
	dc.Stroke()

	dc.DrawCircle(x3, y3, 3)
	// vykreslení
	dc.Stroke()
}

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

	// vykreslení cest
	drawCubicBezier(dc, 50, 220, 50, 50, 160, 220, 160, 50)
	drawCubicBezier(dc, 270, 220, 320, 100, 270, 100, 320, 220)
}
