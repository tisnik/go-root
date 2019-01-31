// Seriál "Programovací jazyk Go"
//
// Desátá část
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
