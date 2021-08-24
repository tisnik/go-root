// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šestá část
//    Konstrukce pro řízení běhu programu v jazyce Go (dokončení)
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go-dokonceni/
//
// Demonstrační příklad číslo 22:
//    Dělení a zbytek po dělení.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_06/22_div_by_zero.html

package main

import "fmt"

func computeDivMod(x, y int) {
	fmt.Printf("%3d / %2d = %3d   %3d %% %2d = %3d\n", x, y, x/y, x, y, x%y)
}

func main() {
	computeDivMod(10, 0)
}
