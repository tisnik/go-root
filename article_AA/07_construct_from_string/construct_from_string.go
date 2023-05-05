package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	d0, err := decimal.NewFromString("-0")
	fmt.Println(d0, err)
	fmt.Println()

	d1, err := decimal.NewFromString("-1234567890.123456789")
	fmt.Println(d1, err)
	fmt.Println()

	d2, err := decimal.NewFromString("1e1000")
	fmt.Println(d2, err)
	fmt.Println()

	d3, err := decimal.NewFromString("1e-1000")
	fmt.Println(d3, err)
	fmt.Println()

	d4, err := decimal.NewFromString("-1234567890e123456")
	fmt.Println(d4, err)
	fmt.Println()
}
