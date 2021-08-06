// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// První část:
//    Go: minimalistický a překvapivě výkonný programovací jazyk
//    https://www.root.cz/clanky/go-minimalisticky-a-prekvapive-vykonny-programovaci-jazyk/
//
// Demonstrační příklad číslo 15:
//    Funkce se dvěma parametry a dvěma návratovými hodnotami
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_01/15_swap_function_B.html

package main

import "fmt"

func swap(a int, b int) (int, int) {
	return b, a
}

func main() {
	x := 1
	y := 2
	z, w := swap(x, y)

	fmt.Println("z =", z)
	fmt.Println("w =", w)
}
