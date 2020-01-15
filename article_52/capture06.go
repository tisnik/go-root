package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
)

func CaptureErrorOutput(function func()) (string, error) {
	// backup of the real stderr
	stderr := os.Stderr

	// temporary replacement for stdout
	reader, writer, err := os.Pipe()
	if err != nil {
		return "", err
	}

	// temporarily replace real Stderr by the mocked one
	defer func() {
		os.Stderr = stderr
	}()
	os.Stderr = writer

	// channel with captured standard output
	captured := make(chan string)

	// synchronization object
	wg := new(sync.WaitGroup)
	// we are going to wait for one goroutine only
	wg.Add(1)

	go func() {
		var buf bytes.Buffer
		// goroutine is started -> inform main one via WaitGroup object
		wg.Done()
		io.Copy(&buf, reader)
		captured <- buf.String()
	}()
	// wait for goroutine to start
	wg.Wait()
	// provided function that (probably) prints something to standard output
	function()
	writer.Close()
	return <-captured, nil
}

func main() {
	str, err := CaptureErrorOutput(func() { fmt.Fprintln(os.Stderr, "Error output again") })
	if err != nil {
		panic(err)
	}

	fmt.Println("Captured error output:")
	fmt.Println("-------------------------------")
	fmt.Println(str)
	fmt.Println("-------------------------------")
}
