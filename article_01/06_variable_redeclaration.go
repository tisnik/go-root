// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// První část:
//    Go: minimalistický a překvapivě výkonný programovací jazyk
//    https://www.root.cz/clanky/go-minimalisticky-a-prekvapive-vykonny-programovaci-jazyk/
//
// Seznam příkladů z první části:
//    https://github.com/tisnik/go-root/blob/master/article_01/README.md
//
// Demonstrační příklad číslo 6:
//    Pokus o opětovnou deklaraci proměnných
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_01/06_variable_redeclaration.html

package main

import "fmt"

func main() {
	a := 10
	fmt.Println(a)
	b := "hello"
	fmt.Println(b)
	c := true
	fmt.Println(c)

	a := 20
	fmt.Println(a)
	b := "world"
	fmt.Println(b)
	c := false
	fmt.Println(c)
}
