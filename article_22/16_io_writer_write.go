// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá druhá část
//     Vstupně-výstupní funkce standardní knihovny programovacího jazyka Go
//     https://www.root.cz/clanky/vstupne-vystupni-funkce-standardni-knihovny-programovaciho-jazyka-go/
//
// Seznam příkladů z dvacáté druhé části:
//    https://github.com/tisnik/go-root/blob/master/article_22/README.md
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
