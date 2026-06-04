// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá třetí část
//     Pokročilejší použití vstupně-výstupních funkcí standardní knihovny jazyka Go
//     https://www.root.cz/clanky/pokrocilejsi-pouziti-vstupne-vystupnich-funkci-standardni-knihovny-jazyka-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z dvacáté třetí části:
//    https://github.com/tisnik/go-root/blob/master/article_23/README.md
//
// Demonstrační příklad číslo 16:
//    Komprimace dat metodou Deflate.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_23/01_io_reader_read.html

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

	gzipWriter := gzip.NewWriter(writer)
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
