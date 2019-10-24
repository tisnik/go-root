// Seriál "Programovací jazyk Go"
//
// Třicátá čtvrtá část
//
// Demonstrační příklad číslo 10:
//    Implementace programové smyčky typu "for" v jazyce Go.

package main

import "fmt"

func main() {
	for x := 1; x <= 10; x++ {
		fmt.Printf("%d\n", x)
	}
}
