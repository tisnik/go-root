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
	v1 := mat.NewVecDense(5, nil)
	v2 := mat.NewVecDense(5, []float64{1, 0, 2, 0, 3})
	v := mat.NewVecDense(5, nil)

	v.SubVec(v1, v2)

	fmt.Printf("v1:\n%v\n\n", mat.Formatted(v1))
	fmt.Printf("v2:\n%v\n\n", mat.Formatted(v2))
	fmt.Printf("v:\n%v\n\n", mat.Formatted(v))
}
