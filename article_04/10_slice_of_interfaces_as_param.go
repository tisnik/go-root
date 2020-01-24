// Seriál "Programovací jazyk Go"
//
// Čtvrtá část
//    Rozhraní, metody, gorutiny a kanály v programovacím jazyku Go
//    https://www.root.cz/clanky/rozhrani-metody-gorutiny-a-kanaly-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 10:
//    Řez s objekty implementující rozhraní předané jako parametr.

package main

import (
	"fmt"
	"math"
)

// ClosedShape je uživatelsky definovaná datová struktura
// představující uzavřené geometrické tvary (úsečka, oblouk, křivka)
type ClosedShape interface {
	area() float64
}

func area(shape ClosedShape) float64 {
	return shape.area()
}

// Circle je uživatelsky definovaná datová struktura
// představující kružnici se středem v bodě [x, y]
// a poloměrem radius
type Circle struct {
	x, y   float64
	radius float64
}

// Ellipse je uživatelsky definovaná datová struktura
// představující elipsu se středem v bodě [x, y]
// a poloměrem poloos a a b
type Ellipse struct {
	x, y float64
	a, b float64
}

// Rectangle je uživatelsky definovaná datová struktura
// představující geometrický tvar obdélníka
type Rectangle struct {
	x, y          float64
	width, height float64
}

func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func (ellipse Ellipse) area() float64 {
	return math.Pi * ellipse.a * ellipse.b
}

func printAreas(shapes []ClosedShape) {
	for _, shape := range shapes {
		fmt.Println(shape)
		fmt.Println(area(shape))
		fmt.Println(shape.area())
		fmt.Println()
	}
}

func main() {
	shapes := []ClosedShape{
		Rectangle{x: 0, y: 0, width: 100, height: 100},
		Circle{x: 0, y: 0, radius: 100},
		Ellipse{x: 0, y: 0, a: 100, b: 50}}

	printAreas(shapes)
}
