package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func factorial(n decimal.Decimal) decimal.Decimal {
	one := decimal.NewFromInt32(1)
	var zero decimal.Decimal

	if n.Cmp(zero) <= 0 {
		return one
	}
	return n.Mul(factorial(n.Sub(one)))
}

func main() {
	for n := int64(1); n < 80; n++ {
		f := factorial(decimal.NewFromInt(n))
		fmt.Printf("%3d! = %s\n", n, f)
	}
}
