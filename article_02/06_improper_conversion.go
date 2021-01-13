// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Druhá část:
//    Datové typy v programovacím jazyku Go
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 6:
//    Přiřazení uint16 -> uint8

package main

import "fmt"

func main() {
	var a uint16 = 255
	var b uint8 = a

	fmt.Println(a)
	fmt.Println(b)
}
