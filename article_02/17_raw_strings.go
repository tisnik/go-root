// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Druhá část:
//    Datové typy v programovacím jazyku Go
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 17
//    Datový typ řetězec

package main

import "fmt"

func main() {
	var s1 string = "Hello\nworld!\n"
	var s2 string = `Hello\nworld!\n`

	fmt.Println(s1)
	fmt.Println(s2)
}
