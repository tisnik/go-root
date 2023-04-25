package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewFloat(1.0)
	y := big.NewFloat(0.0)

	var result big.Float
	result.Quo(x, y)

	fmt.Println(result.Text('f', 31))

	fmt.Println()

	fmt.Println(x.IsInf())
	fmt.Println(y.IsInf())
	fmt.Println(result.IsInf())
}
