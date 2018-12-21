// Seriál "Programovací jazyk Go"
//
// Šestá část
//
// Demonstrační příklad číslo 4:
//    Špatné použití příkazu goto.

package main

import "fmt"

func main() {
	goto Next
	i := 10
Next:
	fmt.Printf("%2d\n", i)
}
