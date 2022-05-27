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
// Demonstrační příklad číslo 3:
//    Základní funkce pro parsing řetězců: ParseInt.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_09/03_string_parsing_int.html

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
