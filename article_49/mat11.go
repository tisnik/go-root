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
	v := mat.NewVecDense(10, nil)
	fmt.Printf("Value:\n%v\n", mat.Formatted(v))
	fmt.Printf("Length:     %d\n", v.Len())
	fmt.Printf("Capacity:   %d\n", v.Cap())
	dim1, dim2 := v.Dims()
	fmt.Printf("Dimensions: %dx%d\n", dim1, dim2)
}
