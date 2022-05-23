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
// Demonstrační příklad číslo 23
//    Kopie řezů polí
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_02/23_slice_copy.html

package main

import "fmt"

func main() {
	var a [10]int

	slice := a[:]

	fmt.Printf("Pole před modifikací: %v\n", a)
	fmt.Printf("Řez před modifikací:  %v\n", slice)

	for i := 0; i < len(a); i++ {
		a[i] = i * 2
	}

	fmt.Printf("Pole po modifikací:   %v\n", a)
	fmt.Printf("Řez po modifikaci:    %v\n", slice)

	for i := 0; i < len(slice); i++ {
		slice[i] = 42
	}

	fmt.Printf("Pole po modifikací:   %v\n", a)
	fmt.Printf("Řez po modifikaci:    %v\n", slice)
}
