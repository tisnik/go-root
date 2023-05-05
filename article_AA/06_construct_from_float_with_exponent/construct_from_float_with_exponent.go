package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	d0 := decimal.NewFromFloatWithExponent(123.456, 0)
	fmt.Println(d0)

	d1 := decimal.NewFromFloatWithExponent(123.456, 1)
	fmt.Println(d1)

	d2 := decimal.NewFromFloatWithExponent(123.456, 2)
	fmt.Println(d2)

	d3 := decimal.NewFromFloatWithExponent(123.456, 3)
	fmt.Println(d3)

	d4 := decimal.NewFromFloatWithExponent(123.456, -1)
	fmt.Println(d4)

	d5 := decimal.NewFromFloatWithExponent(123.456, -2)
	fmt.Println(d5)

	d6 := decimal.NewFromFloatWithExponent(123.456, -3)
	fmt.Println(d6)

	d7 := decimal.NewFromFloatWithExponent(123.456, -4)
	fmt.Println(d7)
}
