// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třetí část
//    Datové typy v programovacím jazyku Go (2.část)
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go-2-cast/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze třetí části:
//    https://github.com/tisnik/go-root/blob/master/article_03/README.md
//
// Demonstrační příklad číslo 28:
//    Kontroly při překladu.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_03/28_improper_func_type.html

package main

import "fmt"

func funkce1(x int) int {
	return 2 * x
}

func funkce2(x, y int) int {
	return x * y
}

func main() {
	var a func(int) int
	fmt.Println(a)

	a = funkce1
	fmt.Println(a)
	fmt.Println(a(10))

	a = funkce2
	fmt.Println(a)
	fmt.Println(a(10, 20))
}
