package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewFloat(2.0)
	y := big.NewFloat(0.1)

	for i := 0; i < 20; i++ {
		value, accuracy := x.Float64()
		fmt.Println(accuracy, value)

		x.Mul(x, y)
	}
}
