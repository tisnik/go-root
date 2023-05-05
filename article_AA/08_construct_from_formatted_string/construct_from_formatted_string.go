package main

import (
	"fmt"
	"regexp"

	"github.com/shopspring/decimal"
)

func main() {
	r := regexp.MustCompile("neco")

	d0, err := decimal.NewFromFormattedString("-0", r)
	fmt.Println(d0, err)
	fmt.Println()

	d1, err := decimal.NewFromFormattedString("-1234567890.123456789", r)
	fmt.Println(d1, err)
	fmt.Println()

	d2, err := decimal.NewFromFormattedString("1e1000", r)
	fmt.Println(d2, err)
	fmt.Println()

	d3, err := decimal.NewFromFormattedString("1e-1000", r)
	fmt.Println(d3, err)
	fmt.Println()

	d4, err := decimal.NewFromFormattedString("-1234567890e123", r)
	fmt.Println(d4, err)
	fmt.Println()
}
