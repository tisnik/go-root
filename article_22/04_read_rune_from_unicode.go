package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("* ěščř ½µ§я¤ *")
	for {
		c, size, err := reader.ReadRune()
		if err == io.EOF {
			fmt.Println("\nEOF detected")
			break
		}
		if err == nil {
			fmt.Printf("%c %d\n", c, size)
		} else {
			fmt.Printf("\nerror %v", err)
			break
		}
	}
}
