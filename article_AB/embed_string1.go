package main

import (
	"embed"
	"fmt"
)

//go:embed hello.txt
var helloMessage string

func main() {
	// pouziti retezce
	fmt.Println(helloMessage)
}
