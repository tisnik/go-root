// Seriál "Programovací jazyk Go"
//
// Dvacátá třetí část
//
// Demonstrační příklad číslo 4:
//     Vylepšení předchozího příkladu.

package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	reader, writer := io.Pipe()

	go func() {
		fmt.Fprint(writer, "Hello Mario!")
		writer.Close()
	}()

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(reader)

	fmt.Printf("Message read from pipe: '%s'\n", buffer.String())
	writer.Close()
}
