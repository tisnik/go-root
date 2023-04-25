package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewFloat(2.0)
	y := big.NewFloat(0.1)

	for i := 0; i < 10; i++ {
		var ratio big.Rat
		_, accuracy := x.Rat(&ratio)
		fmt.Println(accuracy, ratio.String())

		x.Mul(x, y)
	}
}
