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
//    https://tisnik.github.io/go-root/article_50/glot04.html

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
