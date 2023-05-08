package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	x := decimal.NewFromInt32(1)

	y := decimal.NewFromInt32(7)

	for precision := 30; precision > 0; precision-- {
		r := x.DivRound(y, int32(precision))
		fmt.Println(r)
	}
}
