package main

import "github.com/fogleman/gg"

func main() {
	const width = 320
	const height = 240

	dc := gg.NewContext(width, height)

	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Clear()

	dc.SavePNG("02.png")
}
