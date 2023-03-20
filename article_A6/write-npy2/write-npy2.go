package main

import (
	"log"
	"os"

	"github.com/sbinet/npyio"
)

func main() {
	f, err := os.Create("int_vector.npy")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatalf("error closing file: %v\n", err)
		}
	}()

	m := []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	err = npyio.Write(f, m)
	if err != nil {
		log.Fatalf("error writing to file: %v\n", err)
	}
}
