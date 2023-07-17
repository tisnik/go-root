package main

import (
	"embed"
	"log"
	"os"
)

//go:embed npe.jpg
var f embed.FS

func main() {
	data, err := f.ReadFile("npe.jpg")
	if err != nil {
		log.Fatal(err)
	}

	// open output file
	fout, err := os.Create("npe2.jpg")
	if err != nil {
		log.Fatal(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		err := fout.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	fout.Write(data)
}
