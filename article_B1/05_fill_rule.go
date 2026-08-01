package main

import (
	"fmt"
	"math"

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

	customColor := gg.Hex("44aadd")
	dc.SetFillBrush(gg.Solid(customColor))

	dc.SetFillRule(gg.FillRuleEvenOdd)

	radius := float64(height) / 3

	w := float64(width)
	h := float64(height)

	// začátek vykreslování cesty
	dc.MoveTo(w/2, h/2-radius)
	for i := range 4 {
		angle := float64(i+1) * 3.0 * 360 / 5.0
		x := radius * math.Sin(angle*math.Pi/180.0)
		y := radius * math.Cos(angle*math.Pi/180.0)
		dc.LineTo(w/2+x, h/2-y)
	}
	dc.ClosePath()

	// vyplnění obrysu
	dc.Fill()

}
