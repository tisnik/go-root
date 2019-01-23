// Seriál "Programovací jazyk Go"
//
// Devátá část
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
