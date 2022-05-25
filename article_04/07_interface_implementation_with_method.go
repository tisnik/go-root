// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtvrtá část
//    Rozhraní, metody, gorutiny a kanály v programovacím jazyku Go
//    https://www.root.cz/clanky/rozhrani-metody-gorutiny-a-kanaly-v-programovacim-jazyku-go/
//
// Seznam příkladů ze čtvrté části:
//    https://github.com/tisnik/go-root/blob/master/article_04/README.md
//
// Demonstrační příklad číslo 7:
//    Implementace jednoduchého rozhraní nazvaného OpenShape metodou.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_04/07_interface_implementation_with_method.html

package main

import (
	"fmt"
	"math"
)

// OpenShape je uživatelsky definovaná datová struktura
// představující otevřené tvary (úsečka, oblouk, křivka)
type OpenShape interface {
	length() float64
}

// Line je uživatelsky definovaná datová struktura
type Line struct {
	x1, y1 float64
	x2, y2 float64
}

func length(shape OpenShape) float64 {
	return shape.length()
}

func (line Line) length() float64 {
	return math.Hypot(line.x1-line.x2, line.y1-line.y2)
}

func main() {
	line1 := Line{x1: 0, y1: 0, x2: 100, y2: 100}

	fmt.Println(line1)

	lineLength := length(line1)
	fmt.Println(lineLength)
}
