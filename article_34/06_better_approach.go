// Seriál "Programovací jazyk Go"
//
// Třicátá čtvrtá část
//
// Demonstrační příklad číslo 6:
//     Rozvětvení s využitím switch namísto podmínky typu "if-then-else".

package main

import "fmt"

func main() {
	x := 10

	switch {
	case x > 0:
		fmt.Println("x is positive number")
	case x == 0:
		fmt.Println("x is zero")
	default:
		fmt.Println("x is negative number")
	}
}
