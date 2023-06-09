// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá část
//    Tvorba grafů v jazyce Go
//    https://www.root.cz/clanky/tvorba-grafu-v-jazyce-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z padesáté části:
//    https://github.com/tisnik/go-root/blob/master/article_50/README.md
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_50/glot07.html

package main

import (
	"github.com/Arafatk/glot"
)

type Plot struct {
	*glot.Plot
}

func (plot *Plot) Save(filename string) {
	plot.Cmd("set terminal pngcairo")
	plot.Cmd("set output '" + filename + "'")
	plot.Cmd("replot")
}

func NewPlot(dimensions int) *Plot {
	plot, err := glot.NewPlot(dimensions, false, false)
	if err != nil {
		panic(err)
	}
	return &Plot{plot}

}

func main() {
	plot := NewPlot(2)
	defer plot.Close()

	plot.AddPointGroup("Measured data", "lines", []float32{1.2, 1.9, 3.9, 8.2, 9.5, 8.1, 4.1, 1.8, 1.5})
	plot.AddPointGroup("Expected data", "lines", []float32{1, 2, 4, 8, 9, 8, 4, 2, 1})

	plot.SetTitle("Plot #7")
	plot.SetXLabel("t")
	plot.SetYLabel("m/s")

	plot.SetXrange(-2, 10)
	plot.SetYrange(0, 10)
	plot.Save("glot07.png")
}
