// Seriál "Programovací jazyk Go"
//
// Devátá část
//     Užitečné balíčky pro každodenní použití jazyka Go
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-jazyka-go/
//
// Demonstrační příklad číslo 3:
//    Základní funkce pro parsing řetězců: ParseInt.

package main

import (
	"fmt"
	"strconv"
)

func tryToParseInteger(s string, base int) {
	i, err := strconv.ParseInt(s, base, 32)
	if err == nil {
		fmt.Printf("%d\n", i)
	} else {
		fmt.Printf("%v\n", err)
	}
}

func main() {
	tryToParseInteger("42", 10)
	tryToParseInteger("42", 0)
	tryToParseInteger("42", 16)
	tryToParseInteger("42", 2)
	tryToParseInteger("42x", 10)
	println()
	tryToParseInteger("-42", 10)
	tryToParseInteger("-42", 0)
	tryToParseInteger("-42", 16)
	tryToParseInteger("-42", 2)
	tryToParseInteger("-42x", 10)
	println()
}
