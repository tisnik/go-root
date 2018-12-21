// Seriál "Programovací jazyk Go"
//
// Šestá část
//
// Demonstrační příklad číslo 5:
//    Špatné použití příkazu goto.

package main

import "fmt"

func main() {
	i := 10
	goto IntoIf
	if i > 0 {
	IntoIf:
		fmt.Printf("%2d\n", i)
	}
}
