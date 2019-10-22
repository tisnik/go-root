// Seriál "Programovací jazyk Go"
//
// Třicátá čtvrtá část
//
// Demonstrační příklad číslo 6:
//    Vnořené podmínky typu "if".

package main

import "fmt"

func main() {
	x := 10

	if x > 0 {
		fmt.Println("x is positive number")
	} else if x == 0 {
		fmt.Println("x is zero")
	} else {
		fmt.Println("x is negative number")
	}
}
