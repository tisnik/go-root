// Seriál "Programovací jazyk Go"
//
// Devátá část
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
