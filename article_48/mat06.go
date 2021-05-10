// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá osmá část
//    Gophernotes: kombinace interaktivního prostředí Jupyteru s jazykem Go
//    https://www.root.cz/clanky/gophernotes-kombinace-interaktivniho-prostredi-jupyteru-s-jazykem-go/
//
// Demonstrační příklad číslo 6:
//     Vytištění obsahu rozsáhlé matice funkcí mat.Formatted

package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	big := mat.NewDense(100, 100, nil)
	for i := 0; i < 100; i++ {
		big.Set(i, i, 1)
	}

	for i := 1; i <= 5; i++ {
		fmt.Println(mat.Formatted(big, mat.Prefix(" "), mat.Excerpt(i)))
		fmt.Println()
	}
}
