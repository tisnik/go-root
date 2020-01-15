// Seriál "Programovací jazyk Go"
//
// Třetí část
//    Datové typy v programovacím jazyku Go (2.část)
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go-2-cast/
//
// Demonstrační příklad číslo 15:
//    Mapa bez inicializace.

package main

import "fmt"

func main() {
	var m1 map[int]string
	fmt.Println(m1)

	m1[0] = "nula"
}
