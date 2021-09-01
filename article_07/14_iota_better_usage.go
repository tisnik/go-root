// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmá část
//    Programovací jazyk Go: dokončení popisu vlastností samotného jazyka
//    https://www.root.cz/clanky/programovaci-jazyk-go-dokonceni-popisu-vlastnosti-samotneho-jazyka/
//
// Demonstrační příklad číslo 14:
//    Konstanty a identifikátor iota.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_07/14_iota_better_usage.html

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
