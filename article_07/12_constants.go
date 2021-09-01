// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmá část
//    Programovací jazyk Go: dokončení popisu vlastností samotného jazyka
//    https://www.root.cz/clanky/programovaci-jazyk-go-dokonceni-popisu-vlastnosti-samotneho-jazyka/
//
// Demonstrační příklad číslo 12:
//    Konstanty.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_07/12_constants.html

package main

import "fmt"

const (
	Pi float64 = 3.1415927
	E          = 2.71828

	z0 int = 0
	z1     = 0
	z2     = z0 + z1
)

func main() {
	fmt.Printf("Pi = %f\n", Pi)
	fmt.Printf("e = %f\n", E)

	fmt.Printf("z0 = %d\n", z0)
	fmt.Printf("z1 = %d\n", z1)

	fmt.Printf("z2 = %d\n", z2)
}
