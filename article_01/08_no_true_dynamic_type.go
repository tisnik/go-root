// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// První část:
//    Go: minimalistický a překvapivě výkonný programovací jazyk
//    https://www.root.cz/clanky/go-minimalisticky-a-prekvapive-vykonny-programovaci-jazyk/
//
// Demonstrační příklad číslo 8:
//    Špatný typ hodnoty přiřazený do proměnných
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_01/08_no_true_dynamic_type.html

package main

import "fmt"

func main() {
	a := 10
	fmt.Println(a)
	b := "hello"
	fmt.Println(b)
	c := true
	fmt.Println(c)

	a = "world"
	fmt.Println(a)
	b = 0
	fmt.Println(b)
	c = nil
	fmt.Println(c)
}
