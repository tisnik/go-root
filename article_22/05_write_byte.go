package main

import (
	"bytes"
	"fmt"
)

func main() {
	buffer := bytes.Buffer{}
	buffer.WriteByte(65)

	b, _ := buffer.ReadByte()
	fmt.Printf("%02x\n", b)

	buffer.WriteByte(0xff)

	b, _ = buffer.ReadByte()
	fmt.Printf("%02x\n", b)
}
