package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const filename = "test_input_no_eoln.txt"
const buffer_size = 16

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
