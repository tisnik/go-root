package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	x := decimal.NewFromInt32(2)

	y := decimal.NewFromInt32(3)

	r1 := x.Add(y)
	r2 := x.Sub(y)
	r3 := x.Mul(y)
	r4 := x.Div(y)
	r5 := x.Mod(y)

	fmt.Println("x   ", x)
	fmt.Println("y   ", y)
	fmt.Println("x+y ", r1)
	fmt.Println("x-y ", r2)
	fmt.Println("x*y ", r3)
	fmt.Println("x/y ", r4)
	fmt.Println("x%y ", r5)
}
