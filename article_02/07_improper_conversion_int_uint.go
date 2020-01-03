// Seriál "Programovací jazyk Go"
//
// Druhá část:
//    Datové typy v programovacím jazyku Go
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 7:
//    Přiřazení uint8 -> int8 a naopak

package main

import "fmt"

func main() {
	var a int8 = 100
	var b uint8 = a

	fmt.Println(a)
	fmt.Println(b)

	var c uint8 = 100
	var d int8 = c

	fmt.Println(c)
	fmt.Println(d)
}
