package main

import (
	"fmt"

	"github.com/gogpu/gg"
	_ "github.com/gogpu/gg/gpu"
	"github.com/gogpu/gg/integration/ggcanvas"
	"github.com/gogpu/gogpu"
	"github.com/gogpu/gogpu/input"
)

type Point struct {
	X float64
	Y float64
}

func drawPath(dc *gg.Context,
	points []Point, currentPoint Point) {

	dc.SetRGBA(0.0, 0.0, 0.0, 1.0)

	for i := 0; i < len(points)-1; i++ {
		dc.DrawLine(
			points[i].X, points[i].Y,
			points[i+1].X, points[i+1].Y,
		)
	}
	dc.Stroke()

	// barva vykreslení
	dc.SetRGBA(0.8, 0.0, 0.0, 1.0)
	dc.DrawCircle(currentPoint.X, currentPoint.Y, 4)

	// kreslení gumové čáry
	if len(points) > 0 {
		lastPointIndex := len(points) - 1
		dc.DrawLine(
			points[lastPointIndex].X, points[lastPointIndex].Y,
			currentPoint.X, currentPoint.Y)
		dc.Stroke()
	}
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

	// vybrané koncové body
	points := []Point{}

	// aktuální bod
	currentPoint := Point{X: 0, Y: 0}

	// callback metoda zavolaná ve chvíli, kdy se má překreslit okno
	app.OnDraw(func(dc *gogpu.Context) {
		if canvas == nil {
			provider := app.GPUContextProvider()
			w, h := dc.Width(), dc.Height()
			canvas, _ = ggcanvas.New(provider, w, h)
		}

		cw, ch := canvas.Size()

		canvas.Draw(func(cc *gg.Context) {
			renderFrame(cc, cw, ch, points, currentPoint)
		})

		canvas.Render(dc.RenderTarget())
	})

	// callback metoda volaná v případě, že se má scéna překreslit
	app.OnUpdate(func(dt float64) {
		inp := app.Input()

		// reakce na stisk levého tlačítka myši
		if inp.Mouse().Pressed(input.MouseButtonLeft) {
			x, y := inp.Mouse().Position()
			point := Point{X: float64(x), Y: float64(y)}
			points = append(points, point)
			// vynucení překreslení okna
			app.RequestRedraw()
		} else {
			x, y := inp.Mouse().Position()
			currentPoint.X = float64(x)
			currentPoint.Y = float64(y)
			// vynucení překreslení okna
			app.RequestRedraw()
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

func renderFrame(dc *gg.Context, width, height int, points []Point, currentPoint Point) {
	// barva pozadí
	dc.SetRGB(0.9, 0.9, 0.7)

	// obdélník s plochou odpovídající oknu
	dc.DrawRectangle(0, 0, float64(width), float64(height))
	dc.Fill()

	// vykreslení cest
	drawPath(dc, points, currentPoint)
}
