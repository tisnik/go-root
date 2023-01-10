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
// Demonstrační příklad číslo 26:
//    Datový typ "funkce"
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_03/26_func_type.html

package main

import "fmt"

func funkce1() {
	fmt.Println("funkce1")
}

func funkce2() {
	fmt.Println("funkce2")
}

func main() {
	var a func()
	fmt.Println(a)

	a = funkce1
	fmt.Println(a)
	a()

	a = funkce2
	fmt.Println(a)
	a()
}
