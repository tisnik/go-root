// Seriál "Programovací jazyk Go"
//
// Dvacátá třetí část
//     Pokročilejší použití vstupně-výstupních funkcí standardní knihovny jazyka Go
//     https://www.root.cz/clanky/pokrocilejsi-pouziti-vstupne-vystupnich-funkci-standardni-knihovny-jazyka-go/
//
// Demonstrační příklad číslo 8:
//     Použití operace Seek pro relativní posun v souboru.

package main

import (
	"fmt"
	"io"
	"strings"
)

const input_string = "****** Hello world! ******[END]"
const bufferSize = 6

func main() {
	reader := strings.NewReader(input_string)

	buffer := make([]byte, bufferSize)

	for {
		read, err := reader.Read(buffer)

		if read > 0 {
			fmt.Printf("read %d bytes translated into '%s'\n", read, buffer[:read])
			reader.Seek(7, io.SeekCurrent)
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
