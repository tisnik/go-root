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
// Demonstrační příklad číslo 4:
//     Vylepšení předchozího příkladu.

package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	reader, writer := io.Pipe()

	go func() {
		fmt.Fprint(writer, "Hello Mario!")
		writer.Close()
	}()

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(reader)

	fmt.Printf("Message read from pipe: '%s'\n", buffer.String())
	writer.Close()
}
