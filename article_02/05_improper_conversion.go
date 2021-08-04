// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Druhá část:
//    Datové typy v programovacím jazyku Go
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 5:
//    Přiřazení uint8 -> uint16
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_02/05_improper_conversion.html

package main

import "fmt"

func main() {
	var a uint8 = 255
	var b uint16 = a

	fmt.Println(a)
	fmt.Println(b)
}
