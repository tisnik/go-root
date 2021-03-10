// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá třetí část
//     Pokročilejší použití vstupně-výstupních funkcí standardní knihovny jazyka Go
//     https://www.root.cz/clanky/pokrocilejsi-pouziti-vstupne-vystupnich-funkci-standardni-knihovny-jazyka-go/
//
// Demonstrační příklad číslo 19:
//     Souběžná komprimace dat více metodami.

package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const filename1 = "test_output.txt.1.gz"
const filename2 = "test_output.txt.2.gz"
const filename3 = "test_output.txt.3.gz"
const filename4 = "test_output.txt.4.gz"

func fileWriter(filename string) io.WriteCloser {
	writer, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	return writer
}

func gzipWriter(writer io.Writer, compressionLevel int) io.WriteCloser {
	gzipWriter, err := gzip.NewWriterLevel(writer, compressionLevel)
	if err != nil {
		log.Fatal(err)
	}
	return gzipWriter
}

func writeMessage(writer io.Writer, message string) {
	buffer := []byte(message)
	written, err := writer.Write(buffer)

	if written >= 0 {
		fmt.Printf("written %d bytes\n", written)
	}

	if err != nil {
		fmt.Printf("I/O error %v\n", err)
	}
}

func main() {
	message := strings.Repeat("Hello ", 100)

	writer1 := fileWriter(filename1)
	defer writer1.Close()

	writer2 := fileWriter(filename2)
	defer writer2.Close()

	writer3 := fileWriter(filename3)
	defer writer3.Close()

	writer4 := fileWriter(filename4)
	defer writer4.Close()

	gzipWriter1 := gzipWriter(writer1, gzip.BestCompression)
	defer gzipWriter1.Close()

	gzipWriter2 := gzipWriter(writer2, gzip.BestSpeed)
	defer gzipWriter2.Close()

	gzipWriter3 := gzipWriter(writer3, gzip.HuffmanOnly)
	defer gzipWriter3.Close()

	gzipWriter4 := gzipWriter(writer4, gzip.NoCompression)
	defer gzipWriter4.Close()

	writer := io.MultiWriter(gzipWriter1, gzipWriter2, gzipWriter3, gzipWriter4)

	writeMessage(writer, message)
}
