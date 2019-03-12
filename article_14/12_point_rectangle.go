// Seriál "Programovací jazyk Go"
//
// Čtrnáctá část
//
// Demonstrační příklad číslo 12:
//     Datové struktury Point a Rectangle

package main

import (
	"image"
)

func main() {
	point1 := image.Pt(10, 10)
	point2 := image.Point{10, 10}

	rectangle1 := image.Rect(0, 0, 320, 240)
	rectangle2 := image.Rectangle{image.Point{0, 0}, image.Point{320, 240}}

	println(point1.String())
	println(point2.String())

	println(rectangle1.String())
	println(rectangle2.String())
}
