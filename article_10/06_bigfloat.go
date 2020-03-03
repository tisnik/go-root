// Seriál "Programovací jazyk Go"
//
// Desátá část
//     Užitečné balíčky pro každodenní použití Go (2), porovnání výkonnosti Go s céčkem
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-go-2-porovnani-vykonnosti-go-s-ceckem/
//
// Demonstrační příklad číslo 6:
//    Použití balíčku big, typ Float.

package main

import (
	"fmt"
	. "math/big"
)

func main() {
	x := NewFloat(1.0)
	y := NewFloat(0.5)

	for i := 1; i < 82; i++ {
		fmt.Println(x.Text('f', 80))
		x.Mul(x, y)
	}
}
