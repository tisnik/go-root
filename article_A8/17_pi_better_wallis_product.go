package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	var result big.Rat

	result.SetFrac64(2, 1)

	one := big.NewInt(1)
	limit := big.NewInt(1000)

	for n := big.NewInt(1); n.Cmp(limit) <= 0; n.Add(n, one) {
		m := big.NewInt(4)
		m.Mul(m, n)
		m.Mul(m, n)
		mn := big.NewInt(0)
		mn.Sub(m, one)

		var item big.Rat
		item.SetFrac(m, mn)
		result.Mul(&result, &item)

		f, _ := result.Float64()
		absError := math.Pi - f
		relError := 100.0 * absError / math.Pi
		fmt.Println(f, "\t", absError, "\t", relError, "\t", result.String())
	}
}
