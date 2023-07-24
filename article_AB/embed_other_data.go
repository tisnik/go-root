package main

import (
	_ "embed"
	"fmt"
)

type Foo struct {
	x string
	y bool
}

//go:embed lorem_ipsum.txt
var loremIpsum Foo

func main() {
	// pouziti retezce
	fmt.Println(loremIpsum)
}
