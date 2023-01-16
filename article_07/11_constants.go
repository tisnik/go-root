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
// Demonstrační příklad číslo 11:
//    Konstanty.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_07/11_constants.html

package main

import "fmt"

const Pi float64 = 3.1415927
const E = 2.71828

const z0 int = 0
const z1 = 0

const z2 = z0 + z1

func main() {
	fmt.Printf("Pi = %f\n", Pi)
	fmt.Printf("e = %f\n", E)

	fmt.Printf("z0 = %d\n", z0)
	fmt.Printf("z1 = %d\n", z1)

	fmt.Printf("z2 = %d\n", z2)
}
