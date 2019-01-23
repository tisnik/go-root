// Seriál "Programovací jazyk Go"
//
// Devátá část
//
// Demonstrační příklad číslo 7:
//    Základní funkce pro formátování řetězců: Itoa.

package main

import (
	. "strconv"
)

func main() {
	println(Itoa(42))
	println(Itoa(0))
	println(Itoa(-1))
}
