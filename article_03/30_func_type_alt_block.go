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
// Demonstrační příklad číslo 30:
//    Datový typ "funkce s návratovou hodnotou"
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_03/30_func_type_alt_block.html

package main

import "fmt"

type (
	noParamFunction     func() int
	twoIntParamFunction func(int, int) int
)

func funkce1() int {
	return 42
}

func funkce2() int {
	return -1
}

func funkce3(x int, y int) int {
	return x + y
}

func funkce4(x int, y int) int {
	return x * y
}

func main() {
	var a noParamFunction
	var b twoIntParamFunction

	fmt.Println(a)
	fmt.Println(b)

	a = funkce1
	fmt.Println(a)
	fmt.Println(a())

	a = funkce2
	fmt.Println(a)
	fmt.Println(a())

	b = funkce3
	fmt.Println(b)
	fmt.Println(b(10, 20))

	b = funkce4
	fmt.Println(b)
	fmt.Println(b(10, 20))
}
