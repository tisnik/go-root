// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá osmá část
//    Gophernotes: kombinace interaktivního prostředí Jupyteru s jazykem Go
//    https://www.root.cz/clanky/gophernotes-kombinace-interaktivniho-prostredi-jupyteru-s-jazykem-go/
//
// Demonstrační příklad číslo 10:
//     Výpočet inverzní matice.

package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func printMatrix(m mat.Matrix, name string) {
	fmt.Printf(" *** %s ***\n", name)
	fmt.Println(mat.Formatted(m))
	fmt.Println()

}

func main() {
	var c mat.Dense

	m2 := mat.NewDense(3, 3, []float64{5, 0, 0, 0, 2, 0, 0, 0, 1})

	printMatrix(m2, "m2")

	c.Inverse(m2)
	printMatrix(&c, "c")

	var d mat.Dense
	d.Mul(&c, m2)
	printMatrix(&d, "d")
}
