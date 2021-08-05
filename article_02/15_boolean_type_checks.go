// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Druhá část:
//    Datové typy v programovacím jazyku Go
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 15
//    Datový typ Boolean - kontroly při překladu
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_02/15_boolean_type_checks.html

package main

import "fmt"

func main() {
	var a bool = true
	var b bool = false
	var c bool = 0
	var d bool = ""

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
