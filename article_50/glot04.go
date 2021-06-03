// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá část
//    Tvorba grafů v jazyce Go
//    https://www.root.cz/clanky/tvorba-grafu-v-jazyce-go/

package main

import (
	"github.com/Arafatk/glot"
)

func main() {
	plot, err := glot.NewPlot(2, false, false)
	if err != nil {
		panic(err)
	}

	plot.AddPointGroup("Measured data", "lines", []float32{1.2, 1.9, 3.9, 8.2, 9.5, 8.1, 4.1, 1.8, 1.5})
	plot.AddPointGroup("Expected data", "lines", []float32{1, 2, 4, 8, 9, 8, 4, 2, 1})

	plot.SetXrange(-2, 10)
	plot.SetYrange(0, 10)
	plot.SavePlot("glot04.png")
}
