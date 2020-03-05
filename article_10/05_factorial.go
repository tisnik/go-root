// Seriál "Programovací jazyk Go"
//
// Desátá část
//     Užitečné balíčky pro každodenní použití Go (2), porovnání výkonnosti Go s céčkem
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-go-2-porovnani-vykonnosti-go-s-ceckem/
//
// Demonstrační příklad číslo 5:
//    Použití balíčku big pro výpočet faktoriálu.

package main

import (
	"fmt"
	"math/big"
)

func factorial(n *big.Int) *big.Int {
	one := big.NewInt(1)
	if n.Cmp(big.NewInt(0)) <= 0 {
		return one
	} else {
		return one.Mul(n, factorial(one.Sub(n, one)))
	}
}

func main() {
	for n := int64(1); n < 80; n++ {
		f := factorial(big.NewInt(n))
		fmt.Printf("%3d! = %s\n", n, f.Text(10))
	}
}
