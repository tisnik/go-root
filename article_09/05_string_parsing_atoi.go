// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Devátá část
//     Užitečné balíčky pro každodenní použití jazyka Go
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-jazyka-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z deváté části:
//    https://github.com/tisnik/go-root/blob/master/article_09/README.md
//
// Demonstrační příklad číslo 5:
//    Základní funkce pro parsing řetězců: Atoi.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_09/05_string_parsing_atoi.html

package main

import (
	"fmt"
	"strconv"
)

func tryToParseInteger(s string) {
	i, err := strconv.Atoi(s)
	if err == nil {
		fmt.Printf("%d\n", i)
	} else {
		fmt.Printf("%v\n", err)
	}
}

func main() {
	tryToParseInteger("42")
	tryToParseInteger("42x")
	tryToParseInteger("")
	println()
	tryToParseInteger("-42")
	tryToParseInteger("-42x")
	tryToParseInteger("-")
	println()
}
