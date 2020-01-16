// Seriál "Programovací jazyk Go"
//
// Třetí část
//    Datové typy v programovacím jazyku Go (2.část)
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go-2-cast/
//
// Demonstrační příklad číslo 22:
//    Základní práce s ukazateli.

package main

import "fmt"

func main() {
	var i int = 42

	fmt.Println(i)

	var pI *int
	fmt.Println(pI)

	pI = &i

	fmt.Println(pI)
	fmt.Println(*pI)

	*pI++

	fmt.Println(i)
	fmt.Println(*pI)
}
