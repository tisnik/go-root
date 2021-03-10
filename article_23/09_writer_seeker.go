// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá třetí část
//     Pokročilejší použití vstupně-výstupních funkcí standardní knihovny jazyka Go
//     https://www.root.cz/clanky/pokrocilejsi-pouziti-vstupne-vystupnich-funkci-standardni-knihovny-jazyka-go/
//
// Demonstrační příklad číslo 9:
//     Použití operace Seek při zápisu do souboru.

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const filename = "small.bin"

func writeMark(writer io.Writer) {
	buffer := []byte("**")
	written, err := writer.Write(buffer)

	if written > 0 {
		fmt.Printf("written %d bytes\n", written)
	}

	if err != nil {
		fmt.Printf("I/O error %v\n", err)
	}
}

func main() {
	writer, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	writeMark(writer)
	writer.Seek(100, io.SeekCurrent)
	writeMark(writer)
	writer.Seek(100, io.SeekCurrent)
	writeMark(writer)
}
