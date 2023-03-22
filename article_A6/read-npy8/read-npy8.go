package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sbinet/npyio"
)

func main() {
	f, err := os.Open("large_vector.npy")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatalf("error closing file: %v\n", err)
		}
	}()

	r, err := npyio.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("npy-header: %v\n", r.Header)
	shape := r.Header.Descr.Shape

	v := make([]byte, shape[0])

	err = r.Read(&v)
	if err != nil {
		log.Fatalf("error reading from file: %v\n", err)
	}

	fmt.Printf("loaded vector size = %v\n", len(v))
}
