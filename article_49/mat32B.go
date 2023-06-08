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
	t1 := mat.NewTriDense(3, mat.Upper, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Printf("Value:\n%v\n\n", mat.Formatted(t1))
	t1.SetTri(0, 2, 100)
	fmt.Printf("New value:\n%v\n\n", mat.Formatted(t1))

	fmt.Println()

	t2 := mat.NewTriDense(3, mat.Lower, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Printf("Old value:\n%v\n\n", mat.Formatted(t2))
	t2.SetTri(2, 0, 100)
	fmt.Printf("New value:\n%v\n\n", mat.Formatted(t2))
}
