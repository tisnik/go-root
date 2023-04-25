package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewFloat(1.0)
	y := big.NewFloat(0.5)

	for i := 1; i < 82; i++ {
		fmt.Println(x.Text('f', 80))
		x.Mul(x, y)
	}
}
