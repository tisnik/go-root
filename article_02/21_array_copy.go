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
// Demonstrační příklad číslo 21
//    Kopie polí
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_02/21_array_copy.html

package main

import "fmt"

func main() {
	var a1 [10]int

	a2 := a1

	fmt.Printf("Pole 1: %v\n", a1)
	fmt.Printf("Pole 2: %v\n", a2)

	for i := 0; i < len(a1); i++ {
		a1[i] = i * 2
	}

	fmt.Printf("Pole 1: %v\n", a1)
	fmt.Printf("Pole 2: %v\n", a2)
}
