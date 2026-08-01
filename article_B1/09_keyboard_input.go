package main

import (
	"fmt"
	"math"

	"github.com/gogpu/gg"
	_ "github.com/gogpu/gg/gpu"
	"github.com/gogpu/gg/integration/ggcanvas"
	"github.com/gogpu/gogpu"
	"github.com/gogpu/gogpu/input"
)

func main() {
	// rozměry okna
	const width = 400
	const height = 300

	rotation := 0.0

	// inicializace aplikace s GUI
	app := gogpu.NewApp(gogpu.DefaultConfig().
		WithTitle("GoGPU + gg").
		WithSize(width, height))

	var canvas *ggcanvas.Canvas
	var red = 0.3
	var green = 0.7
	var blue = 0.9

	// callback metoda zavolaná ve chvíli, kdy se má překreslit okno
	app.OnDraw(func(dc *gogpu.Context) {
		if canvas == nil {
			provider := app.GPUContextProvider()
			w, h := dc.Width(), dc.Height()
			canvas, _ = ggcanvas.New(provider, w, h)
		}

		cw, ch := canvas.Size()

		canvas.Draw(func(cc *gg.Context) {
			renderFrame(cc, cw, ch, rotation, red, green, blue)
		})

		canvas.Render(dc.RenderTarget())
	})

	// callback metoda volaná v případě, že se má scéna překreslit
	app.OnUpdate(func(dt float64) {
		inp := app.Input()
		const rotationDelta = 1.5
		const colorDelta = 0.02

		keyboard := inp.Keyboard()
		switch {
		// reakce na stisk šipky doleva
		case keyboard.Pressed(input.KeyLeft):
			rotation -= rotationDelta
			// vynucení překreslení okna
			app.RequestRedraw()

			// reakce na stisk šipky doprava
		case keyboard.Pressed(input.KeyRight):
			rotation += rotationDelta
			// vynucení překreslení okna
			app.RequestRedraw()

			// reakce na stisk klávesy Q
		case keyboard.Pressed(input.KeyQ):
			// změna červené barvové složky
			red = math.Min(1.0, red+colorDelta)
			// vynucení překreslení okna
			app.RequestRedraw()

		// reakce na stisk klávesy W
		case keyboard.Pressed(input.KeyW):
			// změna zelené barvové složky
			green = math.Min(1.0, green+colorDelta)
			// vynucení překreslení okna
			app.RequestRedraw()

			// reakce na stisk klávesy E
		case keyboard.Pressed(input.KeyE):
			// změna modré barvové složky
			blue = math.Min(1.0, blue+colorDelta)
			// vynucení překreslení okna
			app.RequestRedraw()

			// reakce na stisk klávesy A
		case keyboard.Pressed(input.KeyA):
			// změna červené barvové složky
			red = math.Max(0.0, red-colorDelta)
			// vynucení překreslení okna
			app.RequestRedraw()

		// reakce na stisk klávesy S
		case keyboard.Pressed(input.KeyS):
			// změna zelené barvové složky
			green = math.Max(0.0, green-colorDelta)
			// vynucení překreslení okna
			app.RequestRedraw()

			// reakce na stisk klávesy D
		case keyboard.Pressed(input.KeyD):
			// změna modré barvové složky
			blue = math.Max(0.0, blue-colorDelta)
			// vynucení překreslení okna
			app.RequestRedraw()

			// reakce na stisk klávesy Escape
		case inp.Keyboard().Pressed(input.KeyEscape):
			app.Quit()
		}

	})

	// callback metoda zavolaná ve chvíli, kdy se okno uzavírá
	app.OnClose(func() {
		fmt.Println("Closing")
		gg.CloseAccelerator()
	})

	// spuštění aplikace
	app.Run()
}

func rad(angle float64) float64 {
	return angle * math.Pi / 180.0
}

func renderFrame(dc *gg.Context, width, height int, rotation, red, green, blue float64) {
	// barva pozadí
	dc.SetRGB(0.9, 0.9, 0.7)

	// obdélník s plochou odpovídající oknu
	dc.DrawRectangle(0, 0, float64(width), float64(height))
	dc.Fill()

	// výpočet barvy hvězdy
	customColor := gg.RGB(red, green, blue)
	dc.SetFillBrush(gg.Solid(customColor))

	dc.SetFillRule(gg.FillRuleEvenOdd)

	radius := float64(height) / 3

	w := float64(width)
	h := float64(height)

	// začátek vykreslování cesty
	dc.MoveTo(
		w/2+radius*math.Sin(rad(rotation)),
		h/2-radius*math.Cos(rad(rotation)))

	// další segmenty cesty
	for i := range 4 {
		angle := float64(i+1)*3.0*360/5.0 + rotation
		x := radius * math.Sin(rad(angle))
		y := radius * math.Cos(rad(angle))
		dc.LineTo(w/2+x, h/2-y)
	}
	// uzavření cesty
	dc.ClosePath()

	// vyplnění obrysu
	dc.Fill()

}
