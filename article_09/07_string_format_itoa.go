// Seriál "Programovací jazyk Go"
//
// Devátá část
//     Užitečné balíčky pro každodenní použití jazyka Go
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-jazyka-go/
//
// Demonstrační příklad číslo 7:
//    Základní funkce pro formátování řetězců: Itoa.

package main

import (
	"strconv"
)

func main() {
	println(strconv.Itoa(42))
	println(strconv.Itoa(0))
	println(strconv.Itoa(-1))
}
