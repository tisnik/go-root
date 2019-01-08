// Seriál "Programovací jazyk Go"
//
// Sedmá část
//
// Demonstrační příklad číslo 13:
//    Konstanty a identifikátor iota.

package main

import "fmt"

const (
	Pondeli = iota
	Utery   = iota
	Streda  = iota
	Ctvrtek = iota
	Patek   = iota
	Sobota  = iota
	Nedele  = iota
)

func main() {
	fmt.Printf("%d\n", Pondeli)
	fmt.Printf("%d\n", Streda)
	fmt.Printf("%d\n", Patek)
}
