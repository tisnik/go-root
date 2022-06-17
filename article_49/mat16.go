// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá devátá část
//    Popis vybraných balíčků nabízených projektem Gonum
//    https://www.root.cz/clanky/popis-vybranych-balicku-nabizenych-projektem-gonum/
//
// Seznam příkladů ze čtyřicáté deváté části:
//    https://github.com/tisnik/go-root/blob/master/article_49/README.md

package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	v2 := mat.NewVecDense(5, []float64{1, 0, 2, 0, 3})
	v := mat.NewVecDense(5, nil)

	v.ScaleVec(10.0, v2)

	fmt.Printf("v2:\n%v\n\n", mat.Formatted(v2))
	fmt.Printf("v:\n%v\n\n", mat.Formatted(v))
}
