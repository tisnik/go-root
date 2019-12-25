// Seriál "Programovací jazyk Go"
//
// Čtyřicátá osmá část
//
// Demonstrační příklad číslo 1:
//     Vytvoření matice konstruktorem mat.NewDense.

package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	zero := mat.NewDense(5, 6, nil)
	fmt.Printf("%v\n", zero)
}
