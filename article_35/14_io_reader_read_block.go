package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const filename = "test_input.txt"
const buffer_size = 16

func main() {
	fin, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	buffer := make([]byte, buffer_size)

	for {
		read, err := fin.Read(buffer)

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
