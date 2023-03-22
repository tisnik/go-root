package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/sbinet/npyio"
)

func main() {
	f, err := os.Create("large_vector.npy")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatalf("error closing file: %v\n", err)
		}
	}()

	v := make([]byte, math.MaxUint32*2)
	fmt.Println(len(v))

	err = npyio.Write(f, v)
	if err != nil {
		log.Fatalf("error writing to file: %v\n", err)
	}
}
