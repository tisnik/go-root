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
// Demonstrační příklad číslo 13
//    Datové typy reprezentující hodnoty s plovoucí řádovou čárkou
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_02/13_complex_types.html

package main

import "fmt"

func main() {
	var a complex64 = -1.5 + 0i
	var b complex64 = 1.5 + 1000i
	var c complex64 = 1e30 + 1e30i
	var d complex64 = 1i

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	var e complex128 = -1.5 + 0i
	var f complex128 = 1.5 + 1000i
	var g complex128 = 1e300 + 1e300i

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}
