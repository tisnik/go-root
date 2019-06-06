// Seriál "Programovací jazyk Go"
//
// Dvacátá třetí část
//
// Demonstrační příklad číslo 6:
//     Použití operace Seek pro posun počítaný od začátku souboru.

package main

import (
	"fmt"
	"io"
	"strings"
)

const input_string = "*** Hello world! ***"
const buffer_size = 6

func main() {
	reader := strings.NewReader(input_string)

	buffer := make([]byte, buffer_size)
	reader.Seek(4, io.SeekStart)

	for {
		read, err := reader.Read(buffer)

		if read > 0 {
			fmt.Printf("read %d bytes translated into '%s'\n", read, buffer[:read])
		}

		if err == io.EOF {
			fmt.Println("reached end of file")
			break
		}

		if err != nil {
			fmt.Printf("other error %v\n", err)
			break
		}
	}
}
