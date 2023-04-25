package main

import (
	"fmt"
	"log"
	"math/big"
)

func factorial(n *big.Int) *big.Int {
	var zero big.Int
	one := big.NewInt(1)
	if n.Cmp(&zero) <= 0 {
		return one
	}
	return one.Mul(n, factorial(one.Sub(n, one)))
}

func main() {
	for n := int64(1); n < 31; n++ {
		f1 := factorial(big.NewInt(n))

		var f2 big.Int
		f2.MulRange(1, n)

		fmt.Printf("%3d! = %s = %s\n", n, f1.Text(10), f2.Text(10))

		if f1.Cmp(&f2) != 0 {
			log.Panic("Different results detected!")
		}
	}
}
