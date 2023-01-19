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
// Demonstrační příklad číslo 5:
//    Použití balíčku big pro výpočet faktoriálu.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_10/05_factorial.html

package main

import (
	"fmt"
	"math/big"
)

func factorial(n *big.Int) *big.Int {
	one := big.NewInt(1)
	if n.Cmp(big.NewInt(0)) <= 0 {
		return one
	}
	return one.Mul(n, factorial(one.Sub(n, one)))
}

func main() {
	for n := int64(1); n < 80; n++ {
		f := factorial(big.NewInt(n))
		fmt.Printf("%3d! = %s\n", n, f.Text(10))
	}
}
