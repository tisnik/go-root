// Seriál "Programovací jazyk Go"
//
// Třicátá čtvrtá část
//    Programovací jazyk Go pro skalní céčkaře
//    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare/
//
// Demonstrační příklad číslo 13:
//     Smyčky a příkaz break.

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
