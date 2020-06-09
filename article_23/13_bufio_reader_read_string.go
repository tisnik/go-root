// Seriál "Programovací jazyk Go"
//
// Dvacátá třetí část
//     Pokročilejší použití vstupně-výstupních funkcí standardní knihovny jazyka Go
//     https://www.root.cz/clanky/pokrocilejsi-pouziti-vstupne-vystupnich-funkci-standardni-knihovny-jazyka-go/
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
