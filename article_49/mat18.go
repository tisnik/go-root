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
	m := mat.NewDense(3, 3, []float64{1, 0, 0, 0, 1, 0, 0, 0, 1})
	v3 := mat.NewVecDense(3, []float64{2, 3, 4})
	v := mat.NewVecDense(3, nil)
	v.MulVec(m, v3)

	fmt.Printf("v:\n%v\n\n", mat.Formatted(v))
}
