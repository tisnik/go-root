package main

import "github.com/fogleman/gg"

func main() {
	const width = 320
	const height = 240

	dc := gg.NewContext(width, height)

	dc.DrawRectangle(0, 0, width, height)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()

	dc.RotateAbout(gg.Radians(15.0), width/2.0, height/2.0)

	dc.SetRGBA(0.0, 0.0, 0.0, 1)

	dc.DrawLine(32, 20, 288, 20)
	dc.Stroke()

	dc.SetDash(10)
	dc.DrawLine(32, 40, 288, 40)
	dc.Stroke()

	dc.SetDash(10, 10)
	dc.DrawLine(32, 60, 288, 60)
	dc.Stroke()

	dc.SetDash(10, 5)
	dc.DrawLine(32, 80, 288, 80)
	dc.Stroke()

	dc.SetDash(10, 5, 2, 5)
	dc.DrawLine(32, 100, 288, 100)
	dc.Stroke()

	dc.SetLineWidth(4.0)
	dc.SetLineCap(gg.LineCapButt)

	dc.DrawLine(32, 140, 288, 140)
	dc.Stroke()

	dc.SetDash(10)
	dc.DrawLine(32, 160, 288, 160)
	dc.Stroke()

	dc.SetDash(10, 10)
	dc.DrawLine(32, 180, 288, 180)
	dc.Stroke()

	dc.SetDash(10, 5)
	dc.DrawLine(32, 200, 288, 200)
	dc.Stroke()

	dc.SetDash(10, 5, 2, 5)
	dc.DrawLine(32, 220, 288, 220)
	dc.Stroke()

	dc.SavePNG("18.png")
}
