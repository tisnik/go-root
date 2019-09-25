package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const filename = "test_input.txt"

func main() {
	fin, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	buffer := make([]byte, 1)

	for {
		read, err := fin.Read(buffer)

		if read > 0 {
			fmt.Printf("byte from file\n")
			fmt.Println(buffer[0])
		} else {
			fmt.Println("empty block (end of file?)")
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
