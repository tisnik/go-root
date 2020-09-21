// Seriál "Programovací jazyk Go"
//
// Třicátá čtvrtá část
//    Programovací jazyk Go pro skalní céčkaře
//    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare/
//
// Demonstrační příklad číslo 11:
//    Implementace programové smyčky typu "for" v jazyce Go.

package main

import "fmt"

func main() {
	for x := 1; x <= 10000; x <<= 1 {
		fmt.Printf("%d\n", x)
	}
}
