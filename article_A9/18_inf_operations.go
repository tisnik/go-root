package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewFloat(1.0)
	y := big.NewFloat(-1.0)
	z := big.NewFloat(0.0)

	var positiveInfinity big.Float
	positiveInfinity.Quo(x, z)
	fmt.Println(positiveInfinity.Text('f', 31))

	var negativeInfinity big.Float
	negativeInfinity.Quo(y, z)
	fmt.Println(negativeInfinity.Text('f', 31))

	var infMultiplyBy0 big.Float
	infMultiplyBy0.Mul(&positiveInfinity, z)
	fmt.Println(infMultiplyBy0.Text('f', 31))
}
