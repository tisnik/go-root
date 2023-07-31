package main

import (
	"embed"
	"fmt"
	"log"
)

//go:embed a/*
//go:embed b/*
var f embed.FS

func main() {
	entries, err := f.ReadDir("a")
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		fmt.Printf("%-25s  %s\n", entry.Name(), entry.Type())
	}

	fmt.Println("*************************************")

	entries, err = f.ReadDir("b")
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		fmt.Printf("%-25s  %s\n", entry.Name(), entry.Type())
	}
}
