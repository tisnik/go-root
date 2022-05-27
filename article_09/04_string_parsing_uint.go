// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Devátá část
//     Užitečné balíčky pro každodenní použití jazyka Go
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-jazyka-go/
//
// Seznam příkladů z deváté části:
//    https://github.com/tisnik/go-root/blob/master/article_09/README.md
//
// Demonstrační příklad číslo 4:
//    Základní funkce pro parsing řetězců: ParseUInt.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_09/04_string_parsing_uint.html

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
