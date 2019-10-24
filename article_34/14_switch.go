// Seriál "Programovací jazyk Go"
//
// Třicátá čtvrtá část
//
// Demonstrační příklad číslo 14:
//     Rozvětvení v jazyce Go.

package main

import "fmt"

func classify(x int) string {
	switch x {
	case 0:
		return "nula"
	case 2, 4, 6, 8:
		return "sudé číslo"
	case 1, 3, 5, 7, 9:
		return "liché číslo"
	default:
		return "?"
	}
}

func main() {
	for x := 0; x <= 10; x++ {
		fmt.Printf("%d: %s\n", x, classify(x))
	}
}
