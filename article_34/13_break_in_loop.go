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
// Demonstrační příklad číslo 13:
//    Smyčky a příkaz break.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_34/13_break_in_loop.html
//

package main

import "fmt"

func main() {
	x := 1

	for { /* endless loop */
		fmt.Printf("%d\n", x)
		x++
		if x > 10 {
			break
		}
	}
}
