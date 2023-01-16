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
// Demonstrační příklad číslo 2:
//    Logické operátory.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_07/02C_boolean_operators_improper_usage.html

package main

import "fmt"

func main() {
	x := true
	y := false

	fmt.Printf("!%v == %v\n", x, !x)
	fmt.Printf("!%v == %v\n", y, !y)

	fmt.Printf("%v && %v == %v\n", x, y, x && y)
	fmt.Printf("%v || %v == %v\n", x, y, x || y)

	fmt.Printf("%v && %v || %v && %v == %v\n", x, y, true, false, x && y || x && false)
	fmt.Printf("%v && %v || %v && %v == %v\n", x, y, true, false, x && y || x && true)
}
