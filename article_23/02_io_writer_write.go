// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá třetí část
//     Pokročilejší použití vstupně-výstupních funkcí standardní knihovny jazyka Go
//     https://www.root.cz/clanky/pokrocilejsi-pouziti-vstupne-vystupnich-funkci-standardni-knihovny-jazyka-go/
//
// Seznam příkladů z dvacáté třetí části:
//    https://github.com/tisnik/go-root/blob/master/article_23/README.md
//
// Demonstrační příklad číslo 2:
//     Použití metody Write z rozhraní Writer pro zápis do souboru.

package main

import (
	"fmt"
	"log"
	"os"
)

const filename = "test_output.txt"
const message = "Hello world!"

func main() {
	writer, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	buffer := []byte(message)

	written, err := writer.Write(buffer)

	if written > 0 {
		fmt.Printf("written %d bytes\n", written)
	}

	if err != nil {
		fmt.Printf("I/O error %v\n", err)
	}
}
