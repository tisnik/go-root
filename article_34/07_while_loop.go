// Seriál "Programovací jazyk Go"
//
// Třicátá čtvrtá část
//
// Demonstrační příklad číslo 7:
//    Implementace programové smyčky typu "while" v jazyce Go.

package main

import "fmt"

func main() {
	x := 1

	for x <= 10 {
		fmt.Printf("%d\n", x)
		x++
	}
}
