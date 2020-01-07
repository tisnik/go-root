// Seriál "Programovací jazyk Go"
//
// Druhá část:
//    Datové typy v programovacím jazyku Go
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 16
//    Datový typ řetězec

package main

import "fmt"

func main() {
	var s1 string = "www.root.cz"
	var s2 string = ""
	var s3 string = "Hello\nworld!\n"
	var s4 string = "шщэюя"

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}
