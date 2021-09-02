// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Devátá část
//     Užitečné balíčky pro každodenní použití jazyka Go
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-jazyka-go/
//
// Demonstrační příklad číslo 7:
//    Základní funkce pro formátování řetězců: Itoa.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_09/07_string_format_itoa.html

package main

import (
	"strconv"
)

func main() {
	println(strconv.Itoa(42))
	println(strconv.Itoa(0))
	println(strconv.Itoa(-1))
}
