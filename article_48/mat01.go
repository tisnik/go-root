// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá osmá část
//    Gophernotes: kombinace interaktivního prostředí Jupyteru s jazykem Go
//    https://www.root.cz/clanky/gophernotes-kombinace-interaktivniho-prostredi-jupyteru-s-jazykem-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze čtyřicáté osmé části:
//    https://github.com/tisnik/go-root/blob/master/article_48/README.md
//
// Demonstrační příklad číslo 1:
//     Vytvoření matice konstruktorem mat.NewDense.

package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	zero := mat.NewDense(5, 6, nil)
	fmt.Printf("%v\n", zero)
}
