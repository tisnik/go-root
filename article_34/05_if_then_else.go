// Seriál "Programovací jazyk Go"
//
// Třicátá čtvrtá část
//    Programovací jazyk Go pro skalní céčkaře
//    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare/
//
// Demonstrační příklad číslo 5:
//     Podmínka typu "if-then-else" v jazyce Go.

package main

import "fmt"

func main() {
	x := 10

	if x > 0 {
		fmt.Println("x is positive number")
	} else {
		fmt.Println("x is negative number or zero")
	}
}
