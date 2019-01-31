// Seriál "Programovací jazyk Go"
//
// Desátá část
//
// Demonstrační příklad číslo 1:
//    Přečtení argumentů předaných na příkazovém řádku.

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Arguments: %d\n", len(os.Args))

	for index, element := range os.Args {
		fmt.Printf("Argument #%d = %s\n", index, element)
	}
}
