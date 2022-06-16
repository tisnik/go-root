// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá osmá část
//    Gophernotes: kombinace interaktivního prostředí Jupyteru s jazykem Go
//    https://www.root.cz/clanky/gophernotes-kombinace-interaktivniho-prostredi-jupyteru-s-jazykem-go/
//
// Seznam příkladů ze čtyřicáté osmé části:
//    https://github.com/tisnik/go-root/blob/master/article_48/README.md
//
// Demonstrační příklad číslo 7:
//     Transpozice matice.

package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	m1 := mat.NewDense(3, 4, nil)
	m2 := mat.NewDense(3, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})

	fmt.Println(" *** m1 ***")
	fmt.Println(mat.Formatted(m1))
	fmt.Println()

	fmt.Println("  *** m2 ***")
	fmt.Println(mat.Formatted(m2))
	fmt.Println()

	m3 := m2.T()

	fmt.Println("  *** m3 ***")
	fmt.Println(mat.Formatted(m3))
	fmt.Println()
}
