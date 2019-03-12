// Seriál "Programovací jazyk Go"
//
// Čtrnáctá část
//
// Demonstrační příklad číslo 13:
//     Základní operace s obdélníky a body

package main

import (
	"image"
)

func main() {
	point1 := image.Pt(10, 10)
	point2 := image.Pt(1000, 10)

	rectangle1 := image.Rect(0, 0, 200, 200)
	rectangle2 := image.Rect(100, 100, 300, 300)
	rectangle3 := image.Rect(300, 300, 400, 400)

	println(point1.In(rectangle1))
	println(point2.In(rectangle1))

	println(rectangle1.Union(rectangle2).String())
	println(rectangle1.Intersect(rectangle2).String())

	println(rectangle1.Overlaps(rectangle2))
	println(rectangle1.Overlaps(rectangle3))
}
