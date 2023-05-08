package main

import (
	"fmt"
	"math"

	"github.com/shopspring/decimal"
)

func main() {
	result := decimal.NewFromInt32(2)

	one := decimal.NewFromInt32(1)
	limit := decimal.NewFromInt32(200)

	for n := decimal.NewFromInt32(1); n.Cmp(limit) <= 0; n = n.Add(one) {
		m := decimal.NewFromInt32(4)
		m = m.Mul(n)
		m = m.Mul(n)
		mn := m.Sub(one)

		m = m.Div(mn)

		result = result.Mul(m)

		f, _ := result.Float64()
		absError := math.Pi - f
		relError := 100.0 * absError / math.Pi
		fmt.Println(f, "\t", absError, "\t", relError) //, "\t", result.String())
	}
}
