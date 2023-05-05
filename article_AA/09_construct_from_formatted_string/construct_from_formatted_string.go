package main

import (
	"fmt"
	"regexp"

	"github.com/shopspring/decimal"
)

func main() {
	r := regexp.MustCompile(",")

	d0, err := decimal.NewFromFormattedString("1,000", r)
	fmt.Println(d0, err)
	fmt.Println()

	d1, err := decimal.NewFromFormattedString("1,000,000", r)
	fmt.Println(d1, err)
	fmt.Println()

	d2, err := decimal.NewFromFormattedString("1.000", r)
	fmt.Println(d2, err)
	fmt.Println()

	d3, err := decimal.NewFromFormattedString("1,000.000", r)
	fmt.Println(d3, err)
	fmt.Println()
}
