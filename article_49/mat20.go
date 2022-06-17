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
	v1 := mat.NewVecDense(5, nil)
	v2 := mat.NewVecDense(5, []float64{1, 0, 2, 0, 3})

	fmt.Printf("dot(v1, v1): %f\n", mat.Dot(v1, v1))
	fmt.Printf("dot(v1, v2): %f\n", mat.Dot(v1, v2))
	fmt.Printf("dot(v2, v2): %f\n", mat.Dot(v2, v2))
	fmt.Printf("max(v2):     %f\n", mat.Max(v2))
	fmt.Printf("min(v2):     %f\n", mat.Min(v2))
	fmt.Printf("sum(v2):     %f\n", mat.Sum(v2))
}
