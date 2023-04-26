package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewFloat(0.0)

	var divideZeroByZero big.Float
	divideZeroByZero.Quo(x, x)
	fmt.Println(x.Text('f', 31))
}
