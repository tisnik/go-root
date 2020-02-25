// Seriál "Programovací jazyk Go"
//
// Devátá část
//     Užitečné balíčky pro každodenní použití jazyka Go
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-jazyka-go/
//
// Demonstrační příklad číslo 2:
//    Základní funkce pro parsing řetězců: ParseBool.

package main

import (
	"fmt"
	"strconv"
)

func tryToParseBool(s string) {
	i, err := strconv.ParseBool(s)
	if err == nil {
		fmt.Printf("%t\n", i)
	} else {
		fmt.Printf("%v\n", err)
	}
}

func main() {
	tryToParseBool("true")
	tryToParseBool("True")
	tryToParseBool("TRUE")
	tryToParseBool("T")
	tryToParseBool("t")
	tryToParseBool("false")
	tryToParseBool("False")
	tryToParseBool("FALSE")
	tryToParseBool("F")
	tryToParseBool("f")
	tryToParseBool("Foobar")
	tryToParseBool("0")
	tryToParseBool("1")
	tryToParseBool("no")
	tryToParseBool("")
}
