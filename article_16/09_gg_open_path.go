package main

import "github.com/fogleman/gg"

func main() {
	const width = 320
	const height = 240

	dc := gg.NewContext(width, height)

	dc.DrawRectangle(0, 0, width, height)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()

	dc.SetRGBA(0.0, 0.0, 0.0, 1)

	dc.MoveTo(100, 200)
	dc.LineTo(100, 100)
	dc.LineTo(150, 50)
	dc.LineTo(200, 100)
	dc.LineTo(200, 200)
	dc.LineTo(100, 200)

	dc.Stroke()

	dc.SavePNG("09.png")
}
