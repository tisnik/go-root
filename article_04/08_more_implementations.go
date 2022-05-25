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
// Demonstrační příklad číslo 8:
//    Několikanásobná implementace rozhraní.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_04/08_more_implementations.html

package main

import (
	"fmt"
	"math"
)

// OpenShape je uživatelsky definovaná datová struktura
// představující otevřené geometrické tvary (úsečka, oblouk, křivka)
type OpenShape interface {
	length() float64
}

// ClosedShape je uživatelsky definovaná datová struktura
// představující uzavřené geometrické tvary (úsečka, oblouk, křivka)
type ClosedShape interface {
	area() float64
}

func length(shape OpenShape) float64 {
	return shape.length()
}

func area(shape ClosedShape) float64 {
	return shape.area()
}

// Line je uživatelsky definovaná datová struktura
// představující úsečku z bodu [x1, y1] do bodu [x2, y2]
type Line struct {
	x1, y1 float64
	x2, y2 float64
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

func (line Line) length() float64 {
	return math.Hypot(line.x1-line.x2, line.y1-line.y2)
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

func main() {
	line1 := Line{x1: 0, y1: 0, x2: 100, y2: 100}
	fmt.Println("Line")
	fmt.Println(line1)
	fmt.Println(length(line1))
	fmt.Println(line1.length())
	fmt.Println()

	fmt.Println("Rectangle")
	r := Rectangle{x: 0, y: 0, width: 100, height: 100}
	fmt.Println(r)
	fmt.Println(area(r))
	fmt.Println(r.area())
	fmt.Println()

	fmt.Println("Circle")
	c := Circle{x: 0, y: 0, radius: 100}
	fmt.Println(c)
	fmt.Println(area(c))
	fmt.Println(c.area())
	fmt.Println()

	fmt.Println("Ellipse")
	e := Ellipse{x: 0, y: 0, a: 100, b: 50}
	fmt.Println(e)
	fmt.Println(area(e))
	fmt.Println(e.area())
	fmt.Println()
}
