// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtvrtá část
//    Rozhraní, metody, gorutiny a kanály v programovacím jazyku Go
//    https://www.root.cz/clanky/rozhrani-metody-gorutiny-a-kanaly-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 1B:
//    Unikátní jména metod.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_04/01_B_unique_names.html

package main

import (
	"fmt"
	"math"
)

// Point je uživatelsky definovaná datová struktura
type Point struct {
	x1, y1 float64
}

// Line je uživatelsky definovaná datová struktura
type Line struct {
	x1, y1 float64
	x2, y2 float64
}

func length() float64 {
	return 0
}

// Metoda definovaná pro strukturu Point
func (point Point) length() float64 {
	return 0
}

// Metoda definovaná pro strukturu Line
func (line Line) length() float64 {
	return math.Hypot(line.x1-line.x2, line.y1-line.y2)
}

func main() {
	line1 := Line{x1: 0, y1: 0, x2: 100, y2: 100}

	fmt.Println(line1)

	lineLength := line1.length()
	fmt.Println(lineLength)
}
