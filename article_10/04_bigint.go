// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Desátá část
//     Užitečné balíčky pro každodenní použití Go (2), porovnání výkonnosti Go s céčkem
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-go-2-porovnani-vykonnosti-go-s-ceckem/
//
// Demonstrační příklad číslo 4:
//    Použití balíčku big, typ Int.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_10/04_bigint.html

package main

import (
	"fmt"
	"math/big"
)

func main() {
	var x big.Int
	var y big.Int
	x.SetInt64(1)
	y.SetInt64(2)

	for i := 1; i < 200; i++ {
		fmt.Println(x.Text(10))
		x.Mul(&x, &y)
	}
}
