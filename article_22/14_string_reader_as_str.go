// Seriál "Programovací jazyk Go"
//
// Dvacátá druhá část
//
// Demonstrační příklad číslo 14:
//     Načítání ze standardního vstupu.

package main

import (
	"fmt"
	"io"
	"strings"
)

const input_string = "Hello world!"
const buffer_size = 4

func main() {
	r := strings.NewReader(input_string)

	buffer := make([]byte, buffer_size)

	for {
		read, err := r.Read(buffer)

		if read > 0 {
			fmt.Printf("read %d bytes\n", read)
			fmt.Println(string(buffer[:read]))
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
