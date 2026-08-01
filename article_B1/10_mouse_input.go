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

type Point struct {
	X float64
	Y float64
}

func drawCubicBezier(dc *gg.Context,
	p0, p1, p2, p3 Point) {

	dc.SetRGBA(1.0, 0.5, 0.5, 1.0)
	dc.DrawLine(p0.X, p0.Y, p1.X, p1.Y)
	dc.Stroke()
	dc.DrawLine(p1.X, p1.Y, p2.X, p2.Y)
	dc.Stroke()
	dc.DrawLine(p2.X, p2.Y, p3.X, p3.Y)
	dc.Stroke()

	dc.SetRGBA(0.0, 0.0, 0.0, 1.0)
	dc.MoveTo(p0.X, p0.Y)
	dc.CubicTo(p1.X, p1.Y, p2.X, p2.Y, p3.X, p3.Y)
	dc.Stroke()

	dc.SetRGBA(0.2, 0.2, 1.0, 1.0)

	dc.DrawCircle(p0.X, p0.Y, 3)
	dc.Stroke()

	dc.DrawCircle(p1.X, p1.Y, 3)
	dc.Stroke()

	dc.DrawCircle(p2.X, p2.Y, 3)
	// vykreslení
	dc.Stroke()

	dc.DrawCircle(p2.X, p2.Y, 3)
	// vykreslení
	dc.Stroke()

	dc.DrawCircle(p3.X, p3.Y, 3)
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

	// všechny řídicí body obou Bézierových kubik
	controlPoints := []Point{
		Point{50, 220},
		Point{50, 50},
		Point{160, 220},
		Point{160, 50},
		Point{270, 220},
		Point{320, 100},
		Point{270, 100},
		Point{320, 220}}

	// callback metoda zavolaná ve chvíli, kdy se má překreslit okno
	app.OnDraw(func(dc *gogpu.Context) {
		if canvas == nil {
			provider := app.GPUContextProvider()
			w, h := dc.Width(), dc.Height()
			canvas, _ = ggcanvas.New(provider, w, h)
		}

		cw, ch := canvas.Size()

		canvas.Draw(func(cc *gg.Context) {
			renderFrame(cc, cw, ch, controlPoints)
		})

		canvas.Render(dc.RenderTarget())
	})

	// callback metoda volaná v případě, že se má scéna překreslit
	app.OnUpdate(func(dt float64) {
		inp := app.Input()

		// reakce na stisk levého tlačítka myši
		if inp.Mouse().Pressed(input.MouseButtonLeft) {
			x, y := inp.Mouse().Position()
			i := findNearestControlPoint(controlPoints, float64(x), float64(y))
			// nastavení nových souřadnic
			controlPoints[i].X = float64(x)
			controlPoints[i].Y = float64(y)
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

func renderFrame(dc *gg.Context, width, height int, cp []Point) {
	// barva pozadí
	dc.SetRGB(0.9, 0.9, 0.7)

	// obdélník s plochou odpovídající oknu
	dc.DrawRectangle(0, 0, float64(width), float64(height))
	dc.Fill()

	// vykreslení cest
	drawCubicBezier(dc, cp[0], cp[1], cp[2], cp[3])
	drawCubicBezier(dc, cp[4], cp[5], cp[6], cp[7])
}

func findNearestControlPoint(controlPoints []Point, x, y float64) int {
	nearestIndex := 0
	nearestDistance := math.MaxFloat64
	for i, point := range controlPoints {
		distance := math.Hypot(point.X-x, point.Y-y)
		if distance < nearestDistance {
			nearestDistance = distance
			nearestIndex = i
		}
	}
	return nearestIndex
}
