package main

import (
	"compress/gzip"
	"fmt"
	"log"
	"os"
)

const filename = "test_output.txt.gz"
const message = "Hello Hello Hello Hello Hello Hello Hello Hello Hello Hello "

func main() {
	writer, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	gzipWriter, err := gzip.NewWriterLevel(writer, gzip.BestCompression)
	if err != nil {
		log.Fatal(err)
	}

	defer gzipWriter.Close()

	buffer := []byte(message)
	written, err := gzipWriter.Write(buffer)

	if written >= 0 {
		fmt.Printf("written %d bytes\n", written)
	}

	if err != nil {
		fmt.Printf("I/O error %v\n", err)
	}
}
