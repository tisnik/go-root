package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"sync"
)

func CaptureStandardOutput(function func()) (string, error) {
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

func printSinus() {
	epsilon := 1e-6
	for x := 0.0; x <= 2.0*math.Pi+epsilon; x += math.Pi / 6.0 {
		fmt.Printf("sin(%5.2f) = %+5.3f\n", x, math.Sin(x))
	}
}

func main() {
	str, err := CaptureStandardOutput(printSinus)
	if err != nil {
		panic(err)
	}
	fmt.Println("Captured output:")
	fmt.Println("-------------------------------")
	fmt.Println(str)
	fmt.Println("-------------------------------")
}
