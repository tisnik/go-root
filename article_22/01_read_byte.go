// Seriál "Programovací jazyk Go"
//
// Dvacátá druhá část
//
// Demonstrační příklad číslo 1:
//     Postupné načtení bajtů z řetězce.

package main

import (
	"fmt"
	"strings"
)

func main() {
	reader := strings.NewReader("Hello world!")
	for {
		b, err := reader.ReadByte()
		if err == nil {
			fmt.Printf("%c", b)
		} else {
			fmt.Printf("\nerror %v", err)
			break
		}
	}
}
