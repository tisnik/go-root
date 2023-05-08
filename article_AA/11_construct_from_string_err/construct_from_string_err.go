package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	d0, err := decimal.NewFromString("-a")
	fmt.Println(d0, err)
	fmt.Println()

	d1, err := decimal.NewFromString("-x.123456789")
	fmt.Println(d1, err)
	fmt.Println()

	d2, err := decimal.NewFromString("1e100a")
	fmt.Println(d2, err)
	fmt.Println()
}
