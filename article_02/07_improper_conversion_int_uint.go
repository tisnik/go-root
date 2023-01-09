// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Druhá část:
//    Datové typy v programovacím jazyku Go
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z druhé části:
//    https://github.com/tisnik/go-root/blob/master/article_02/README.md
//
// Demonstrační příklad číslo 7:
//    Přiřazení uint8 -> int8 a naopak
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_02/07_improper_conversion_int_uint.html

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
