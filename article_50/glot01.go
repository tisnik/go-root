// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá část
//    Tvorba grafů v jazyce Go
//    https://www.root.cz/clanky/tvorba-grafu-v-jazyce-go/
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_50/glot01.html

package main

import (
	"github.com/Arafatk/glot"
)

func main() {
	plot, err := glot.NewPlot(2, false, false)
	if err != nil {
		panic(err)
	}

	plot.AddPointGroup("Measured data", "points", []int32{1, 2, 4, 8, 9, 8, 4, 2, 1})
	plot.SavePlot("glot01.png")
}
