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
// Demonstrační příklad číslo 4:
//    Podmínka typu "if-then" v jazyce Go.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_34/04_if_then.html
//

package main

import "fmt"

func main() {
	x := 10

	if x > 0 {
		fmt.Println("x is positive number")
	}
}
