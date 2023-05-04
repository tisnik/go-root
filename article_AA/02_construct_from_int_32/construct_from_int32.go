package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	d0 := decimal.NewFromInt32(0)
	fmt.Println(d0)

	d1 := decimal.NewFromInt32(1234)
	fmt.Println(d1)

	d2 := decimal.NewFromInt32(12345)
	fmt.Println(d2)

	d3 := decimal.NewFromInt32(123456)
	fmt.Println(d3)

	d4 := decimal.NewFromInt32(1234567)
	fmt.Println(d4)

	d5 := decimal.NewFromInt32(12345678)
	fmt.Println(d5)
}
