// Seriál "Programovací jazyk Go"
//
// Dvacátá třetí část
//     Pokročilejší použití vstupně-výstupních funkcí standardní knihovny jazyka Go
//     https://www.root.cz/clanky/pokrocilejsi-pouziti-vstupne-vystupnich-funkci-standardni-knihovny-jazyka-go/
//
// Demonstrační příklad číslo 17:
//     Volba metody a úrovně komprimace.

package main

import (
	"compress/gzip"
	"fmt"
	"log"
	"os"
)

const filename = "test_output.txt.gz"
const message = "Hello Hello Hello Hello Hello Hello Hello Hello Hello Hello "

func main() {
	writer, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	gzipWriter, err := gzip.NewWriterLevel(writer, gzip.BestCompression)
	if err != nil {
		log.Fatal(err)
	}

	defer gzipWriter.Close()

	buffer := []byte(message)
	written, err := gzipWriter.Write(buffer)

	if written >= 0 {
		fmt.Printf("written %d bytes\n", written)
	}

	if err != nil {
		fmt.Printf("I/O error %v\n", err)
	}
}
