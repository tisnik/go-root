// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmá část
//    Programovací jazyk Go: dokončení popisu vlastností samotného jazyka
//    https://www.root.cz/clanky/programovaci-jazyk-go-dokonceni-popisu-vlastnosti-samotneho-jazyka/
//
// Demonstrační příklad číslo 2B:
//    Nekorektní použití operátorů na pravdivostní hodnoty.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_07/02B_boolean_operators_improper_usage.html

package main

import "fmt"

func main() {
	x := true
	y := false

	fmt.Printf("^%v == %v\n", x, ^x)
	fmt.Printf("^%v == %v\n", y, ^y)

	fmt.Printf("%v & %v == %v\n", x, y, x&y)
	fmt.Printf("%v | %v == %v\n", x, y, x|y)
	fmt.Printf("%v + %v == %v\n", x, y, x+y)
	fmt.Printf("%v - %v == %v\n", x, y, x-y)
}
