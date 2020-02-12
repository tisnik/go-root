// Seriál "Programovací jazyk Go"
//
// Sedmá část
//    Programovací jazyk Go: dokončení popisu vlastností samotného jazyka
//    https://www.root.cz/clanky/programovaci-jazyk-go-dokonceni-popisu-vlastnosti-samotneho-jazyka/
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
