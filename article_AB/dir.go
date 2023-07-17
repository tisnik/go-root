package main

import (
	"embed"
	"fmt"
	"log"
)

//go:embed *.txt
var f embed.FS

func main() {
	entries, err := f.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		fmt.Printf("%-25s  %s\n", entry.Name(), entry.Type())
	}
}
