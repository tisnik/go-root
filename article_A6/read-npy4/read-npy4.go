package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sbinet/npyio"
)

func main() {
	f, err := os.Open("float_matrix.npy")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var m []float64
	err = npyio.Read(f, &m)
	if err != nil {
		log.Fatalf("error reading from file: %v\n", err)
	}

	fmt.Printf("loaded vector = %v\n", m)

	err = f.Close()
	if err != nil {
		log.Fatalf("error closing file: %v\n", err)
	}
}
