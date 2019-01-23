// Seriál "Programovací jazyk Go"
//
// Devátá část
//
// Demonstrační příklad číslo 5:
//    Základní funkce pro parsing řetězců: Atoi.

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
