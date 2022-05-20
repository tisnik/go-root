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
// Demonstrační příklad číslo 13:
//    Funkce se dvěma parametry
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_01/13_sum_function.html

package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("sum =", sum(1, 2))
}
