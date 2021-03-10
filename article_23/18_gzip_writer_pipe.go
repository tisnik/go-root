// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá třetí část
//     Pokročilejší použití vstupně-výstupních funkcí standardní knihovny jazyka Go
//     https://www.root.cz/clanky/pokrocilejsi-pouziti-vstupne-vystupnich-funkci-standardni-knihovny-jazyka-go/
//
// Demonstrační příklad číslo 18:
//     Přečtení zkomprimovaných dat zpět.

package main

import (
	"compress/gzip"
	"fmt"
	"io"
)

const message = "Hello Hello Hello Hello Hello Hello Hello Hello Hello Hello "

func main() {
	reader, writer := io.Pipe()

	gzipWriter := gzip.NewWriter(writer)

	go func() {
		buffer := []byte(message)
		written, err := gzipWriter.Write(buffer)

		if written >= 0 {
			fmt.Printf("written %d bytes\n", written)
		}

		if err != nil {
			fmt.Printf("I/O error %v\n", err)
		}

		gzipWriter.Close()
		writer.Close()
	}()

	buffer := make([]byte, 100)

	for {
		read, err := reader.Read(buffer)

		if read > 0 {
			fmt.Printf("read %d bytes:", read)
			for i := 0; i < read; i++ {
				fmt.Printf(" %02x", buffer[i])
			}
			fmt.Println()
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
