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
//    https://tisnik.github.io/go-root/article_50/plot01.html

package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
)

const resX = 20.0 / 3.0 * vg.Inch
const resY = 5.0 * vg.Inch

func main() {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	err = p.Save(resX, resY, "plot01.png")
	if err != nil {
		panic(err)
	}
}
