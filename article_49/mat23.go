// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá devátá část
//    Popis vybraných balíčků nabízených projektem Gonum
//    https://www.root.cz/clanky/popis-vybranych-balicku-nabizenych-projektem-gonum/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze čtyřicáté deváté části:
//    https://github.com/tisnik/go-root/blob/master/article_49/README.md

package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	v := mat.NewSymDense(3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	t := v.T()

	fmt.Printf("Value:\n%v\n\n", mat.Formatted(t))

	dim1, dim2 := t.Dims()
	fmt.Printf("Dimensions: %dx%d\n", dim1, dim2)
}
