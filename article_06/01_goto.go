// Seriál "Programovací jazyk Go"
//
// Šestá část
//
// Demonstrační příklad číslo 1:
//    Použití příkazu goto.

package main

import "fmt"

func main() {
	i := 10
Next_i:
	fmt.Printf("%2d\n", i)
	i--
	if i >= 0 {
		goto Next_i
	}
}
