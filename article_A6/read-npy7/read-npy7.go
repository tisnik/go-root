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

	r, err := npyio.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("npy-header: %v\n", r.Header)
	shape := r.Header.Descr.Shape

	v := make([]float64, shape[0]*shape[1])

	err = r.Read(&v)
	if err != nil {
		log.Fatalf("error reading from file: %v\n", err)
	}

	fmt.Printf("loaded vector = %v\n", v)

	m := mat.NewDense(shape[0], shape[1], v)
	fmt.Printf("converted matrix =\n%v\n", mat.Formatted(m))

	err = f.Close()
	if err != nil {
		log.Fatalf("error closing file: %v\n", err)
	}
}
