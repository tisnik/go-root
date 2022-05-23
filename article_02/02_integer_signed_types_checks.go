// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Druhá část:
//    Datové typy v programovacím jazyku Go
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go/
//
// Seznam příkladů z druhé části:
//    https://github.com/tisnik/go-root/blob/master/article_02/README.md
//
// Demonstrační příklad číslo 2:
//    Kontrola celočíselných konstant
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_02/02_integer_signed_types_checks.html

package main

import "fmt"

func main() {
	var a int8 = -1000
	var b int16 = -100000
	var c int32 = -10000000000
	var d int32 = -10000000000000000

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
