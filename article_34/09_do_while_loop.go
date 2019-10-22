// Seriál "Programovací jazyk Go"
//
// Třicátá čtvrtá část
//
// Demonstrační příklad číslo 9:
//    Implementace programové smyčky typu "do-while" v jazyce Go.

package main

import "fmt"

func main() {
	x := 1

	for cond := true; cond; cond = x <= 10 {
		fmt.Printf("%d\n", x)
		x++
	}
}
