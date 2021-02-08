// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmá část
//    Programovací jazyk Go: dokončení popisu vlastností samotného jazyka
//    https://www.root.cz/clanky/programovaci-jazyk-go-dokonceni-popisu-vlastnosti-samotneho-jazyka/
//
// Demonstrační příklad číslo 2:
//    Logické operátory.

package main

import "fmt"

func main() {
	x := 1
	y := 2

	fmt.Printf("!%v == %v\n", x, !x)
	fmt.Printf("!%v == %v\n", y, !y)

	fmt.Printf("%v && %v == %v\n", x, y, x && y)
	fmt.Printf("%v || %v == %v\n", x, y, x || y)

	fmt.Printf("%v && %v || %v && %v == %v\n", x, y, true, false, x && y || x && false)
	fmt.Printf("%v && %v || %v && %v == %v\n", x, y, true, false, x && y || x && true)
}
