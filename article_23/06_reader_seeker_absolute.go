// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá třetí část
//     Pokročilejší použití vstupně-výstupních funkcí standardní knihovny jazyka Go
//     https://www.root.cz/clanky/pokrocilejsi-pouziti-vstupne-vystupnich-funkci-standardni-knihovny-jazyka-go/
//
// Demonstrační příklad číslo 6:
//     Použití operace Seek pro posun počítaný od začátku souboru.

package main

import (
	"fmt"
	"io"
	"strings"
)

const inputString = "*** Hello world! ***"
const bufferSize = 6

func main() {
	reader := strings.NewReader(inputString)

	buffer := make([]byte, bufferSize)
	reader.Seek(4, io.SeekStart)

	for {
		read, err := reader.Read(buffer)

		if read > 0 {
			fmt.Printf("read %d bytes translated into '%s'\n", read, buffer[:read])
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
