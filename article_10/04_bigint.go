// Seriál "Programovací jazyk Go"
//
// Desátá část
//
// Demonstrační příklad číslo 4:
//    Použití balíčku big, typ Int.

package main

import (
	"fmt"
	. "math/big"
)

func main() {
	var x Int
	var y Int
	x.SetInt64(1)
	y.SetInt64(2)

	for i := 1; i < 200; i++ {
		fmt.Println(x.Text(10))
		x.Mul(&x, &y)
	}
}
