package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sbinet/npyio"
	"gonum.org/v1/gonum/mat"
)

func main() {
	f, err := os.Open("float_matrix.npy")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var v []float64
	err = npyio.Read(f, &v)
	if err != nil {
		log.Fatalf("error reading from file: %v\n", err)
	}

	fmt.Printf("loaded vector = %v\n", v)

	m := mat.NewDense(3, 4, v)
	fmt.Printf("converted matrix =\n%v\n", mat.Formatted(m))

	err = f.Close()
	if err != nil {
		log.Fatalf("error closing file: %v\n", err)
	}
}
