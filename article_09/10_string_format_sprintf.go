// Seriál "Programovací jazyk Go"
//
// Devátá část
//     Užitečné balíčky pro každodenní použití jazyka Go
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-jazyka-go/
//
// Demonstrační příklad číslo 10:
//    Základní funkce pro formátování řetězců: Sprintf.

package main

import (
	"fmt"
	"math"
)

func main() {
	value := 42
	s := fmt.Sprintf("%10d", value)
	println(s)

	s = fmt.Sprintf("%10.5f", math.Pi)
	println(s)

	s = fmt.Sprintf("%10.9f", math.Pi)
	println(s)

	a := []int{10, 20, 30}

	s = fmt.Sprintf("%v", a)
	println(s)

}
