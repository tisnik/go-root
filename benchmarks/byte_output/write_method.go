package main

import (
	"bufio"
	"bytes"
	//"io"
	"os"
	"time"
)

func main() {
	length := 100000000
	var buffer bytes.Buffer

	for i := 0; i < length; i++ {
		buffer.WriteString("abc\n")
	}

	data := buffer.String()

	f, _ := os.Create("/tmp/d1")
	defer f.Close()

	t1 := time.Now()
	writer := bufio.NewWriter(f)
	writer.Write([]byte(data))
	writer.Flush()

	elapsed := time.Since(t1)
	println(elapsed)
}
