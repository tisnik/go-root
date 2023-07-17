package main

import (
	_ "embed"
	"fmt"
)

//go:embed lorem_ipsum.txt
var loremIpsum string

func main() {
	// pouziti retezce
	fmt.Println(loremIpsum)
}
