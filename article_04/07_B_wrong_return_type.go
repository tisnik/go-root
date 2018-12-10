// Seriál "Programovací jazyk Go"
//
// Čtvrtá část
//
// Demonstrační příklad číslo 7:
//    Implementace jednoduchého rozhraní nazvaného OpenShape metodou.

package main

import (
	"fmt"
	"math"
)

type OpenShape interface {
	length() float32
}

type Line struct {
	x1, y1 float64
	x2, y2 float64
}

func length(shape OpenShape) float32 {
	return shape.length()
}

func (line Line) length() float64 {
	return math.Hypot(line.x1-line.x2, line.y1-line.y2)
}

func main() {
	l1 := Line{x1: 0, y1: 0, x2: 100, y2: 100}

	fmt.Println(l1)

	line_length := length(l1)
	fmt.Println(line_length)
}
