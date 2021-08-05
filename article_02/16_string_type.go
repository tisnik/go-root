// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Druhá část:
//    Datové typy v programovacím jazyku Go
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 16
//    Datový typ řetězec
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_02/16_string_type.html

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
