// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmá část
//    Programovací jazyk Go: dokončení popisu vlastností samotného jazyka
//    https://www.root.cz/clanky/programovaci-jazyk-go-dokonceni-popisu-vlastnosti-samotneho-jazyka/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze sedmé části:
//    https://github.com/tisnik/go-root/blob/master/article_07/README.md
//
// Demonstrační příklad číslo 13:
//    Konstanty a identifikátor iota.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_07/13_iota.html

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
