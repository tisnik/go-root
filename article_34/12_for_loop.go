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
// Demonstrační příklad číslo 12:
//    Implementace programové smyčky typu "for" v jazyce Go.

package main

import "fmt"

func main() {
	for i, x := 0, 1; x <= 10000; i, x = i+1, x<<1 {
		fmt.Printf("%d\n", x)
	}
}
