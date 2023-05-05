package main

import (
	"fmt"
	"regexp"

	"github.com/shopspring/decimal"
)

func main() {
	r := regexp.MustCompile("\\$")

	d0, err := decimal.NewFromFormattedString("$1000", r)
	fmt.Println(d0, err)
	fmt.Println()

	d1, err := decimal.NewFromFormattedString("20.25$", r)
	fmt.Println(d1, err)
	fmt.Println()

	r = regexp.MustCompile("Kč")
	d2, err := decimal.NewFromFormattedString("1000Kč", r)
	fmt.Println(d2, err)
	fmt.Println()
}
