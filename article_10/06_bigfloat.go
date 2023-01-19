// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Desátá část
//     Užitečné balíčky pro každodenní použití Go (2), porovnání výkonnosti Go s céčkem
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-go-2-porovnani-vykonnosti-go-s-ceckem/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z desáté části:
//    https://github.com/tisnik/go-root/blob/master/article_10/README.md
//
// Demonstrační příklad číslo 6:
//    Použití balíčku big, typ Float.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_10/06_bigfloat.html

package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewFloat(1.0)
	y := big.NewFloat(0.5)

	for i := 1; i < 82; i++ {
		fmt.Println(x.Text('f', 80))
		x.Mul(x, y)
	}
}
