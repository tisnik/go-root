// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá čtvrtá část
//    Programovací jazyk Go pro skalní céčkaře
//    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze třicáté čtvrté části:
//    https://github.com/tisnik/go-root/blob/master/article_34/README.md
//
// Demonstrační příklad číslo 6:
//     Rozvětvení s využitím switch namísto podmínky typu "if-then-else".

package main

import "fmt"

func main() {
	x := 10

	switch {
	case x > 0:
		fmt.Println("x is positive number")
	case x == 0:
		fmt.Println("x is zero")
	default:
		fmt.Println("x is negative number")
	}
}
