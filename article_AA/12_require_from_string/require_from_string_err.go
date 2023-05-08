package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	d0 := decimal.RequireFromString("-a")
	fmt.Println(d0)
	fmt.Println()

	d1 := decimal.RequireFromString("-x.123456789")
	fmt.Println(d1)
	fmt.Println()

	d2 := decimal.RequireFromString("1e100a")
	fmt.Println(d2)
	fmt.Println()
}
