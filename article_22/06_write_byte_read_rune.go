// Seriál "Programovací jazyk Go"
//
// Dvacátá druhá část část
//
// Demonstrační příklad číslo 6:
//     Zápis bajtů du bufferu, přečtení ve formě run.

package main

import (
	"bytes"
	"fmt"
)

func main() {
	buffer := bytes.Buffer{}
	buffer.WriteByte(65)

	c, size, err := buffer.ReadRune()
	fmt.Printf("%c %d %v\n", c, size, err)

	buffer.WriteByte(0xc4)
	buffer.WriteByte(0x9b)

	c, size, err = buffer.ReadRune()
	fmt.Printf("%c %d %v\n", c, size, err)

	c, size, err = buffer.ReadRune()
	fmt.Printf("%c %d %v\n", c, size, err)
}
