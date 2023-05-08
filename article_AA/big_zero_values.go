package main

import (
	"fmt"
	"math/big"
)

func main() {
	var x big.Int
	var y big.Rat
	var z big.Float

	fmt.Println("big.Int  ", x.Text(10))
	fmt.Println("big.Rat  ", y.String())
	fmt.Println("big.Float", z.Text('f', 10))
}
