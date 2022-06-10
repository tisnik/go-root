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
// Demonstrační příklad číslo 7:
//    Implementace programové smyčky typu "while" v jazyce Go.

package main

import "fmt"

func main() {
	x := 1

	for x <= 10 {
		fmt.Printf("%d\n", x)
		x++
	}
}
