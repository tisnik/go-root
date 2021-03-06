// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá devátá část
//    Popis vybraných balíčků nabízených projektem Gonum
//    https://www.root.cz/clanky/popis-vybranych-balicku-nabizenych-projektem-gonum/

package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	d := mat.NewDiagDense(10, nil)
	fmt.Printf("Value:\n%v\n\n", mat.Formatted(d))

	dim1, dim2 := d.Dims()
	fmt.Printf("Dimensions: %dx%d\n", dim1, dim2)
}
