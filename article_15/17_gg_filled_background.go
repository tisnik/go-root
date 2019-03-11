package main

import "github.com/fogleman/gg"

func main() {
	const width = 320
	const height = 240

	dc := gg.NewContext(width, height)

	dc.DrawRectangle(0, 0, width, height)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()

	dc.DrawLine(10, 10, width-10, height-10)
	dc.SetRGB(0.2, 0.0, 0.8)
	dc.Stroke()

	dc.DrawLine(10, height-10, width-10, 10)
	dc.SetRGB(0.8, 0.0, 0.2)
	dc.Stroke()

	dc.SavePNG("17.png")
}
