package main

import (
	"fmt"
	"io"
	"strings"
)

const input_string = "****** Hello world! ******[END]"
const buffer_size = 6

func main() {
	reader := strings.NewReader(input_string)

	buffer := make([]byte, buffer_size)

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
