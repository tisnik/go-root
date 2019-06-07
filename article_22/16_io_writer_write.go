// Seriál "Programovací jazyk Go"
//
// Dvacátá druhá část
//
// Demonstrační příklad číslo 16:
//     Zápis do souboru.

package main

import (
	"fmt"
	"log"
	"os"
)

const filename = "test_output.txt"
const message = "Hello world!"

func main() {
	fout, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fout.Close()

	buffer := []byte(message)
	written, err := fout.Write(buffer)

	if written > 0 {
		fmt.Printf("written %d bytes\n", written)
	}

	if err != nil {
		fmt.Printf("I/O error %v\n", err)
	}
}
