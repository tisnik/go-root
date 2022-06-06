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
// Demonstrační příklad číslo 1:
//     	Čtení dat ze souboru s využitím metody Read z rozhraní Reader.

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const filename = "test_input.txt"
const bufferSize = 16

func main() {
	reader, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	buffer := make([]byte, bufferSize)

	for {
		read, err := reader.Read(buffer)

		if read > 0 {
			fmt.Printf("read %d bytes\n", read)
			fmt.Println(buffer[:read])
		}

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
