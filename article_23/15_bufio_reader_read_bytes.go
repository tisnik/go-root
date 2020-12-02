// Seriál "Programovací jazyk Go"
//
// Dvacátá třetí část
//     Pokročilejší použití vstupně-výstupních funkcí standardní knihovny jazyka Go
//     https://www.root.cz/clanky/pokrocilejsi-pouziti-vstupne-vystupnich-funkci-standardni-knihovny-jazyka-go/
//
// Demonstrační příklad číslo 15:
//     Čtení sekvence bajtů s využitím funkcí z balíčku bufio.

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const filename = "test_input_no_eoln.txt"
const bufferSize = 16

func main() {
	reader, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	bufferedReader := bufio.NewReader(reader)

	for {
		array, err := bufferedReader.ReadBytes('\n')

		fmt.Printf("read array of bytes with size %d bytes: %v\n", len(array), array)

		if err == io.EOF {
			fmt.Println("reached end of file")
			break
		}

		if err != nil {
			fmt.Printf("other error %v\n", err)
			break
		}
	}
}
