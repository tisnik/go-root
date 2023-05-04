package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	d0 := decimal.NewFromInt(0)
	fmt.Println(d0)

	d1 := decimal.NewFromInt(123456)
	fmt.Println(d1)

	d2 := decimal.NewFromInt(123456789)
	fmt.Println(d2)

	d3 := decimal.NewFromInt(12345678900000)
	fmt.Println(d3)

	d4 := decimal.NewFromInt(1234567890000000000)
	fmt.Println(d4)
}
