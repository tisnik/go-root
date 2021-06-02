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
	d := mat.NewDiagDense(10, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Printf("Value:\n%v\n\n", mat.Formatted(d))

	d.SetDiag(1, 100)
	fmt.Printf("Value:\n%v\n\n", mat.Formatted(d))
}
