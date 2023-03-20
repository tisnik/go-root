package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sbinet/npyio"
)

func main() {
	f, err := os.Open("int8_vector.npy")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatalf("error closing file: %v\n", err)
		}
	}()

	var m []int8
	err = npyio.Read(f, &m)
	if err != nil {
		log.Fatalf("error reading from file: %v\n", err)
	}

	fmt.Printf("loaded vector = %v\n", m)
}
