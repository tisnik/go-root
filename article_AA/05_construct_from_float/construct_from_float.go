package main

import (
	"fmt"
	"math"

	"github.com/shopspring/decimal"
)

func main() {
	d0 := decimal.NewFromFloat(0)
	fmt.Println(d0)

	d1 := decimal.NewFromFloat(1e10)
	fmt.Println(d1)

	d2 := decimal.NewFromFloat(1e-10)
	fmt.Println(d2)

	d3 := decimal.NewFromFloat(1.2e308)
	fmt.Println(d3)

	d4 := decimal.NewFromFloat(1.2e-308)
	fmt.Println(d4)

	d5 := decimal.NewFromFloat(1.1e-1000)
	fmt.Println(d5)

	d6 := decimal.NewFromFloat(math.Inf(1))
	fmt.Println(d6)

	d7 := decimal.NewFromFloat(math.Inf(-1))
	fmt.Println(d7)
}
