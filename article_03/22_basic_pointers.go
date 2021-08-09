// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třetí část
//    Datové typy v programovacím jazyku Go (2.část)
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go-2-cast/
//
// Demonstrační příklad číslo 22:
//    Základní práce s ukazateli.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_03/22_basic_pointers.html

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
