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
// Demonstrační příklad číslo 13:
//     Čtení řetězců pomocí funkcí z balíčku bufio.

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const filename = "test_input.txt"

func main() {
	reader, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	bufferedReader := bufio.NewReader(reader)

	for {
		str, err := bufferedReader.ReadString('\n')

		fmt.Printf("read string with size %d bytes: %s", len(str), str)

		if err == io.EOF {
			fmt.Println("\nreached end of file")
			break
		}

		if err != nil {
			fmt.Printf("\nother error %v\n", err)
			break
		}
	}
}
