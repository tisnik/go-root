package main

import (
	"fmt"
	"io"
)

func main() {
	reader, writer := io.Pipe()

	done := make(chan bool)

	go func() {
		fmt.Fprint(writer, "Hello Mario!")
		writer.Close()
		done <- true
	}()

	buffer := make([]byte, 100)

	read, err := reader.Read(buffer)
	if err != nil {
		panic(err)
	} else {
		if read > 0 {
			fmt.Printf("read %d bytes translated into '%s'\n", read, buffer)
		}
	}
	<-done
	reader.Close()
}
