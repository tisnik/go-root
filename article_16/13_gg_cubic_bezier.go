package main

import "github.com/fogleman/gg"

func main() {
	const width = 320
	const height = 240

	dc := gg.NewContext(width, height)

	dc.DrawRectangle(0, 0, width, height)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()

	dc.SetRGBA(1.0, 0.0, 0.0, 1.0)
	dc.MoveTo(10, 180)
	dc.CubicTo(10, 10, 120, 180, 120, 10)
	dc.Stroke()

	dc.SetRGBA(0.0, 1.0, 0.0, 1.0)
	dc.MoveTo(110, 180)
	dc.CubicTo(190, 100, 80, 100, 160, 180)
	dc.Stroke()

	dc.SetRGBA(0.0, 0.0, 1.0, 1.0)
	dc.MoveTo(230, 180)
	dc.CubicTo(280, 60, 230, 60, 280, 180)
	dc.Stroke()

	dc.SavePNG("13.png")
}
