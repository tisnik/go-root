// Seriál "Programovací jazyk Go"
//
// Třicátá čtvrtá část
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
