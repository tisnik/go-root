package capture

import (
	"bytes"
	"io"
	"os"
	"sync"
)

// StandardOutput captures the standard output for specified code
// block and then returns the captured output.
// Please see https://medium.com/@hau12a1/golang-capturing-log-println-and-fmt-println-output-770209c791b4
// for further explanation how it works under the hood.
func StandardOutput(function func()) (string, error) {
	// backup of the real stdout
	stdout := os.Stdout

	// temporary replacement for stdout
	reader, writer, err := os.Pipe()
	if err != nil {
		return "", err
	}

	// temporarily replace real Stdout by the mocked one
	defer func() {
		os.Stdout = stdout
	}()
	os.Stdout = writer

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
