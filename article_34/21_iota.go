// Seriál "Programovací jazyk Go"
//
// Třicátá čtvrtá část
//
// Demonstrační příklad číslo 21:
//    Deklarace konstant v Go.

package main

import "fmt"

const (
	Pondeli = iota
	Utery
	Streda
	Ctvrtek
	Patek
	Sobota
	Nedele
)

func main() {
	fmt.Printf("%d\n", Pondeli)
	fmt.Printf("%d\n", Streda)
	fmt.Printf("%d\n", Patek)
}
