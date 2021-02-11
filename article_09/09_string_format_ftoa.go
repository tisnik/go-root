// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Devátá část
//     Užitečné balíčky pro každodenní použití jazyka Go
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-jazyka-go/
//
// Demonstrační příklad číslo 9:
//    Základní funkce pro formátování řetězců: FormatFloat.

package main

import (
	"math"
	"strconv"
)

func main() {
	value := math.Pi
	println(strconv.FormatFloat(value, 'f', 5, 64))
	println(strconv.FormatFloat(value, 'f', 10, 64))
	println(strconv.FormatFloat(value, 'e', 10, 64))
	println(strconv.FormatFloat(value, 'g', 10, 64))
	println(strconv.FormatFloat(value, 'b', 10, 64))

	println()

	value = 1e20
	println(strconv.FormatFloat(value, 'f', 5, 64))
	println(strconv.FormatFloat(value, 'f', 10, 64))
	println(strconv.FormatFloat(value, 'e', 10, 64))
	println(strconv.FormatFloat(value, 'g', 10, 64))
	println(strconv.FormatFloat(value, 'b', 10, 64))

	println()

	value = 0
	println(strconv.FormatFloat(value, 'f', 5, 64))
	println(strconv.FormatFloat(value, 'f', 10, 64))
	println(strconv.FormatFloat(value, 'e', 10, 64))
	println(strconv.FormatFloat(value, 'g', 10, 64))
	println(strconv.FormatFloat(value, 'b', 5, 64))
}
