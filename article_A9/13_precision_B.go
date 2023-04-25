package main

import (
	"fmt"
	"math/big"
)

func main() {
	result := big.NewFloat(2.0)

	one := big.NewInt(1)
	limit := big.NewInt(200)

	for n := big.NewInt(1); n.Cmp(limit) <= 0; n.Add(n, one) {
		m := big.NewInt(4)
		m.Mul(m, n)
		m.Mul(m, n)
		mn := big.NewInt(0)
		mn.Sub(m, one)

		var item big.Rat
		item.SetFrac(m, mn)

		var itemFloat big.Float
		itemFloat.SetRat(&item)

		result.Mul(result, &itemFloat)

	}

	result.SetPrec(0)
	fmt.Println(result.Text('f', 31))
}
