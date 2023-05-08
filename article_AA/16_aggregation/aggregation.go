package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	x := decimal.NewFromInt32(2)
	y := decimal.NewFromInt32(3)
	z := decimal.NewFromInt32(1)
	w := decimal.NewFromInt32(0)

	min := decimal.Min(x, y, z, w)
	max := decimal.Max(x, y, z, w)
	avg := decimal.Avg(x, y, z, w)
	sum := decimal.Sum(x, y, z, w)

	fmt.Println("min   ", min)
	fmt.Println("max   ", max)
	fmt.Println("avg   ", avg)
	fmt.Println("sum   ", sum)
}
