package main

import (
	"embed"
	"fmt"
	"log"
)

//go:embed lorem_ipsum.txt
var f embed.FS

func main() {
	data, err := f.ReadFile("lorem_ipsum.txt")
	if err != nil {
		log.Fatal(err)
	}

	// pouziti retezce
	fmt.Println(string(data))
}
