// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Devátá část
//     Užitečné balíčky pro každodenní použití jazyka Go
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-jazyka-go/
//
// Demonstrační příklad číslo 6:
//    Základní funkce pro parsing řetězců: parseFloat.

package main

import (
	"fmt"
	"strconv"
)

func tryToParseFloat(s string) {
	i, err := strconv.ParseFloat(s, 32)
	if err == nil {
		fmt.Printf("%f\n", i)
	} else {
		fmt.Printf("%v\n", err)
	}
}

func main() {
	tryToParseFloat("42")
	tryToParseFloat("42.0")
	tryToParseFloat(".5")
	tryToParseFloat("0.5")
	tryToParseFloat("5e10")
	tryToParseFloat(".5e10")
	tryToParseFloat(".5e-5")
	tryToParseFloat("-.5e-5")
	tryToParseFloat("5E10")
	tryToParseFloat(".5E10")
	tryToParseFloat(".5E-5")
	tryToParseFloat("-.5E-5")
}
