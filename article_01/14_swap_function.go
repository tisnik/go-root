// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// První část:
//    Go: minimalistický a překvapivě výkonný programovací jazyk
//    https://www.root.cz/clanky/go-minimalisticky-a-prekvapive-vykonny-programovaci-jazyk/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z první části:
//    https://github.com/tisnik/go-root/blob/master/article_01/README.md
//
// Demonstrační příklad číslo 14:
//    Funkce se dvěma parametry a dvěma návratovými hodnotami
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_01/14_swap_function.html

package main

import "fmt"

func swap(a int, b int) (int, int) {
	return b, a
}

func main() {
	x := 1
	y := 2

	var z int
	var w int

	z, w = swap(x, y)
	fmt.Println("z =", z)
	fmt.Println("w =", w)
}
