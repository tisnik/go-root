// Seriál "Programovací jazyk Go"
//
// Třicátá čtvrtá část
//
// Demonstrační příklad číslo 12:
//    Implementace programové smyčky typu "for" v jazyce Go.

package main

import "fmt"

func main() {
	for i, x := 0, 1; x <= 10000; i, x = i+1, x<<1 {
		fmt.Printf("%d\n", x)
	}
}
