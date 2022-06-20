// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá část
//    Tvorba grafů v jazyce Go
//    https://www.root.cz/clanky/tvorba-grafu-v-jazyce-go/
//
// Seznam příkladů z padesáté části:
//    https://github.com/tisnik/go-root/blob/master/article_50/README.md
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_50/glot16.html

package main

import (
	"github.com/Arafatk/glot"
	"math"
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

const points = 200

func main() {
	plot := NewPlot(3)
	defer plot.Close()

	var pointsX [points]float64
	var pointsY [points]float64
	for i := 0; i < points; i++ {
		t := float64(i) * 8.0 * math.Pi / points
		pointsX[i] = math.Sin(t)
		pointsY[i] = math.Cos(t)
	}

	z := 0.0
	function1 := func(u, v float64) float64 {
		z = z + 1.0
		return z
	}

	plot.AddFunc3d("spiral", "lines", pointsX[:], pointsY[:], function1)

	plot.SetTitle("Plot #16")
	plot.SetXLabel("t")
	plot.SetYLabel("y")

	plot.Save("glot16.png")
}
