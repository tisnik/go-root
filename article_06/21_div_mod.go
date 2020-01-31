// Seriál "Programovací jazyk Go"
//
// Šestá část
//
// Demonstrační příklad číslo 21:
//    Dělení a zbytek po dělení.

package main

import "fmt"

func computeDivMod(x, y int) {
	fmt.Printf("%3d / %2d = %3d   %3d %% %2d = %3d\n", x, y, x/y, x, y, x%y)
}

func main() {
	computeDivMod(10, 3)
	computeDivMod(-10, 3)
	computeDivMod(10, -3)
	computeDivMod(-10, -3)

	fmt.Println()

	for i := 1; i <= 10; i++ {
		computeDivMod(100, i)
	}
}
