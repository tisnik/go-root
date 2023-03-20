package main

import (
	"log"
	"os"

	"github.com/sbinet/npyio"
	"gonum.org/v1/gonum/mat"
)

func main() {
	f, err := os.Create("float_matrix.npy")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatalf("error closing file: %v\n", err)
		}
	}()

	m := mat.NewDense(3, 4, []float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
		10, 11, 12})

	err = npyio.Write(f, m)
	if err != nil {
		log.Fatalf("error writing to file: %v\n", err)
	}
}
