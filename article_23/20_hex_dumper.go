// Seriál "Programovací jazyk Go"
//
// Dvacátá třetí část
//
// Demonstrační příklad číslo 20:
//     Zobrazení libovolných binárních dat ve formě hexadecimálního výpisu.

package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

const filename = "test_output.hex"
const message = "*** Hello world! ***"

func main() {
	writer, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	hexDumper := hex.Dumper(writer)
	defer hexDumper.Close()

	buffer := []byte(message)
	written, err := hexDumper.Write(buffer)

	if written >= 0 {
		fmt.Printf("written %d bytes\n", written)
	}

	if err != nil {
		fmt.Printf("I/O error %v\n", err)
	}
}
