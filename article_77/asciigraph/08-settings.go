// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmdesátá sedmá část:
//    Jazyk Go a textový terminál ve funkci základního prvku uživatelského rozhraní (2.část)
//    https://www.root.cz/clanky/jazyk-go-a-textovy-terminal-ve-funkci-zakladniho-prvku-uzivatelskeho-rozhrani-2-cast/
//
// Seznam příkladů ze sedmdesáté sedmé části:
//    https://github.com/tisnik/go-root/blob/master/article_77/README.md

package main

import (
	"fmt"

	"github.com/guptarohit/asciigraph"
)

const SAMPLES = 64

func displayGraph(modulus int, slope int) {
	var values [SAMPLES]float64

	for i := 0; i < SAMPLES; i++ {
		values[i] = float64((i - SAMPLES/2) / slope % modulus)
	}

	width := asciigraph.Width(40)
	height := asciigraph.Height(20)
	caption := asciigraph.Caption(fmt.Sprintf("Modulus: %d, slope: %d", modulus, slope))

	graph := asciigraph.Plot(values[:], width, height, caption)

	fmt.Println(graph)
}

func main() {
	const modulus = 16
	for slope := 1; slope <= 5; slope++ {
		displayGraph(modulus, slope)
		fmt.Println("\n")
	}
}
