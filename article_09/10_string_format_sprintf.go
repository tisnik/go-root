// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Devátá část
//     Užitečné balíčky pro každodenní použití jazyka Go
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-jazyka-go/
//
// Demonstrační příklad číslo 10:
//    Základní funkce pro formátování řetězců: Sprintf.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_09/10_string_format_sprintf.html

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
