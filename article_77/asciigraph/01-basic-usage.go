// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmdesátá sedmá část:
//    Jazyk Go a textový terminál ve funkci základního prvku uživatelského rozhraní (2.část)
//    https://www.root.cz/clanky/jazyk-go-a-textovy-terminal-ve-funkci-zakladniho-prvku-uzivatelskeho-rozhrani-2-cast/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam příkladů ze sedmdesáté sedmé části:
//    https://github.com/tisnik/go-root/blob/master/article_77/README.md

package main

import (
	"fmt"

	"github.com/guptarohit/asciigraph"
)

func main() {
	data := []float64{3, 4, 5, 6, 9, 8, 5, 8, 6, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)

	fmt.Println(graph)
}
