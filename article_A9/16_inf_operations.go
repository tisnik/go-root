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

	var infMultiplyBy1 big.Float
	infMultiplyBy1.Mul(&positiveInfinity, x)
	fmt.Println(infMultiplyBy1.Text('f', 31))

	var infMultiplyByMinus1 big.Float
	infMultiplyByMinus1.Mul(&positiveInfinity, y)
	fmt.Println(infMultiplyByMinus1.Text('f', 31))

	var infMultiplyByNegativeInf big.Float
	infMultiplyByNegativeInf.Mul(&positiveInfinity, &negativeInfinity)
	fmt.Println(infMultiplyByNegativeInf.Text('f', 31))
}
