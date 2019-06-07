// Seriál "Programovací jazyk Go"
//
// Dvacátá druhá část
//
// Demonstrační příklad číslo 4:
//     Načítání jednotlivých run (ne bajtů) z řetězce.

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("* ěščř ½µ§я¤ *")
	for {
		c, size, err := reader.ReadRune()
		if err == io.EOF {
			fmt.Println("\nEOF detected")
			break
		}
		if err == nil {
			fmt.Printf("%c %d\n", c, size)
		} else {
			fmt.Printf("\nerror %v", err)
			break
		}
	}
}
