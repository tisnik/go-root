// Seriál "Programovací jazyk Go"
//
// Šestá část
//
// Demonstrační příklad číslo 25:
//    Bitové posuny zkombinované s přiřazením.

package main

import "fmt"

func main() {
	x := 1

	for i := uint(0); i <= 10; i++ {
		fmt.Printf("%d << %2d == %4d\n", 1, i, x)
		x <<= 1
	}

	fmt.Println()

	x = 10000000

	for i := uint(0); i <= 10; i++ {
		fmt.Printf("%d >> %2d == %4d\n", 1, i, x)
		x >>= 1
	}

}
