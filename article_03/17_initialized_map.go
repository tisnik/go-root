// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třetí část
//    Datové typy v programovacím jazyku Go (2.část)
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go-2-cast/
//
// Demonstrační příklad číslo 17:
//    Mapa s inicializací.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_03/17_initialized_map.html

package main

import "fmt"

func main() {
	m1 := make(map[int]string)
	fmt.Println(m1)

	m1[0] = "nula"
	m1[1] = "jedna"
	m1[2] = "dva"
	m1[3] = "tri"
	m1[4] = "ctyri"
	m1[5] = "pet"
	m1[6] = "sest"

	fmt.Println(m1)
}
