// Seriál "Programovací jazyk Go"
//
// Čtvrtá část
//    Rozhraní, metody, gorutiny a kanály v programovacím jazyku Go
//    https://www.root.cz/clanky/rozhrani-metody-gorutiny-a-kanaly-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 6:
//    Implementace jednoduchého rozhraní nazvaného OpenShape.

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

// Line je datová struktura, ke které mohou být přidány metody
type Line struct {
	x1, y1 float64
	x2, y2 float64
}

func length(shape OpenShape) float64 {
	return shape.length()
}

func main() {
	line1 := Line{x1: 0, y1: 0, x2: 100, y2: 100}

	fmt.Println(line1)

	lineLength := length(line1)
	fmt.Println(lineLength)
}
