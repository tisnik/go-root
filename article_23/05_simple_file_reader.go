// Seriál "Programovací jazyk Go"
//
// Dvacátá třetí část
//     Pokročilejší použití vstupně-výstupních funkcí standardní knihovny jazyka Go
//     https://www.root.cz/clanky/pokrocilejsi-pouziti-vstupne-vystupnich-funkci-standardni-knihovny-jazyka-go/
//
// Demonstrační příklad číslo 5:
//     Čtení dat z řetězce po blocích.

package main

import (
	"fmt"
	"io"
	"strings"
)

const input_string = "*** Hello world! ***"
const buffer_size = 6

func main() {
	reader := strings.NewReader(input_string)

	buffer := make([]byte, buffer_size)

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
