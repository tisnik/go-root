// Seriál "Programovací jazyk Go"
//
// Dvacátá druhá část
//
// Demonstrační příklad číslo 3:
//     Postupné načtení bajtů z řetězce, ukázka použití Unicode.

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("* ěščř ½µ§я¤ *")
	for {
		b, err := reader.ReadByte()
		if err == io.EOF {
			fmt.Println("\nEOF detected")
			break
		}
		if err == nil {
			fmt.Printf("%02x ", b)
		} else {
			fmt.Printf("\nerror %v", err)
			break
		}
	}
}
