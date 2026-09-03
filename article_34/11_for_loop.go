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
// Demonstrační příklad číslo 11:
//    Implementace programové smyčky typu "for" v jazyce Go.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_34/11_for_loop.html
//

package main

import "fmt"

func main() {
	for x := 1; x <= 10000; x <<= 1 {
		fmt.Printf("%d\n", x)
	}
}
