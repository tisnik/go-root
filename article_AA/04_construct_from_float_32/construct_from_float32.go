package main

import (
	"fmt"
	"math"

	"github.com/shopspring/decimal"
)

func main() {
	d0 := decimal.NewFromFloat32(0)
	fmt.Println(d0)

	d1 := decimal.NewFromFloat32(1e10)
	fmt.Println(d1)

	d2 := decimal.NewFromFloat32(1e-10)
	fmt.Println(d2)

	d3 := decimal.NewFromFloat32(1.2e38)
	fmt.Println(d3)

	d4 := decimal.NewFromFloat32(1.2e-38)
	fmt.Println(d4)

	d5 := decimal.NewFromFloat32(1.1e-100)
	fmt.Println(d5)

	d6 := decimal.NewFromFloat32(float32(math.Inf(1)))
	fmt.Println(d6)

	d7 := decimal.NewFromFloat32(float32(math.Inf(-1)))
	fmt.Println(d7)
}
