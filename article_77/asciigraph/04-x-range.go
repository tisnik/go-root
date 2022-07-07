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

const SAMPLES = 200

func main() {
	var values [SAMPLES]float64

	for i := 0; i < SAMPLES; i++ {
		values[i] = float64(i / 2 % 16)
	}
	graph := asciigraph.Plot(values[:])

	fmt.Println(graph)
}
