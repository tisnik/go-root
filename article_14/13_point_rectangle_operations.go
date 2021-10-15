// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtrnáctá část
//     Programovací jazyk Go a počítačová grafika (úvod)
//     https://www.root.cz/clanky/programovaci-jazyk-go-a-pocitacova-grafika-uvod/
//
// Demonstrační příklad číslo 13:
//     Základní operace s obdélníky a body
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_14/13_point_rectangle_operations.html

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
