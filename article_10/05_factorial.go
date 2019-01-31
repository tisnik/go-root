// Seriál "Programovací jazyk Go"
//
// Desátá část
//
// Demonstrační příklad číslo 5:
//    Použití balíčku big pro výpočet faktoriálu.

package main

import (
	"fmt"
	. "math/big"
)

func factorial(n *Int) *Int {
	one := NewInt(1)
	if n.Cmp(NewInt(0)) <= 0 {
		return one
	} else {
		return one.Mul(n, factorial(one.Sub(n, one)))
	}
}

func main() {
	for n := int64(1); n < 80; n++ {
		f := factorial(NewInt(n))
		fmt.Printf("%3d! = %s\n", n, f.Text(10))
	}
}
