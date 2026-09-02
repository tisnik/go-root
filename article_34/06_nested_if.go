// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá čtvrtá část
//    Programovací jazyk Go pro skalní céčkaře
//    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare/
//
// Seznam příkladů ze třicáté čtvrté části:
//    https://github.com/tisnik/go-root/blob/master/article_34/README.md
//
// Demonstrační příklad číslo 6:
//    Vnořené podmínky typu "if".
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_34/06_nested_if.html
//

package main

import "fmt"

func main() {
	x := 10

	if x > 0 {
		fmt.Println("x is positive number")
	} else if x == 0 {
		fmt.Println("x is zero")
	} else {
		fmt.Println("x is negative number")
	}
}
