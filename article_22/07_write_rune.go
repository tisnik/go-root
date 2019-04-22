package main

import (
	"bytes"
	"fmt"
)

func main() {
	buffer := bytes.Buffer{}
	buffer.WriteRune('a')
	buffer.WriteRune('ě')
	buffer.WriteRune('я')

	for {
		b, err := buffer.ReadByte()
		if err == nil {
			fmt.Printf("%02x ", b)
		} else {
			fmt.Printf("\nerror %v", err)
			break
		}
	}
}
