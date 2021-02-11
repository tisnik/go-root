// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Devátá část
//     Užitečné balíčky pro každodenní použití jazyka Go
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-jazyka-go/
//
// Demonstrační příklad číslo 4:
//    Základní funkce pro parsing řetězců: ParseUInt.

package main

import (
	"fmt"
	"strconv"
)

func tryToParseUnsignedInteger(s string, base int) {
	i, err := strconv.ParseUint(s, base, 32)
	if err == nil {
		fmt.Printf("%d\n", i)
	} else {
		fmt.Printf("%v\n", err)
	}
}

func main() {
	tryToParseUnsignedInteger("42", 10)
	tryToParseUnsignedInteger("42", 0)
	tryToParseUnsignedInteger("42", 16)
	tryToParseUnsignedInteger("42", 2)
	tryToParseUnsignedInteger("42x", 10)
	println()
	tryToParseUnsignedInteger("-42", 10)
	tryToParseUnsignedInteger("-42", 0)
	tryToParseUnsignedInteger("-42", 16)
	tryToParseUnsignedInteger("-42", 2)
	tryToParseUnsignedInteger("-42x", 10)
	println()
}
