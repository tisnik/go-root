// Seriál "Programovací jazyk Go"
//
// Šestá část
//
// Demonstrační příklad číslo 6:
//    Špatné použití příkazu goto.

package main

import "fmt"

func main() {
	goto IntoLoop

	for i := 0; i < 10; i++ {
	IntoLoop:
		fmt.Printf("%2d\n", i)
	}
}
