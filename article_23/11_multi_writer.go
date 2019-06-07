// Seriál "Programovací jazyk Go"
//
// Dvacátá třetí část
//
// Demonstrační příklad číslo 11:
//     Souběžný zápis do více souborů.

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const filename1 = "test_output_1.txt"
const filename2 = "test_output_2.txt"
const message = "Hello world!"

func main() {
	writer1, err := os.Create(filename1)
	if err != nil {
		log.Fatal(err)
	}
	defer writer1.Close()

	writer2, err := os.Create(filename2)
	if err != nil {
		log.Fatal(err)
	}
	defer writer2.Close()

	writer := io.MultiWriter(writer1, writer2)

	buffer := []byte(message)

	written, err := writer.Write(buffer)

	if written > 0 {
		fmt.Printf("written %d bytes\n", written)
	}

	if err != nil {
		fmt.Printf("I/O error %v\n", err)
	}
}
