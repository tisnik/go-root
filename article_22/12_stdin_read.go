package main

import (
	"fmt"
	"io"
	"os"
)

const buffer_size = 16

func main() {
	buffer := make([]byte, buffer_size)

	for {
		read, err := os.Stdin.Read(buffer)

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
