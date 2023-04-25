package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewFloat(1)
	y := big.NewFloat(99)

	for i := 0; i < 40; i++ {
		x.Mul(x, y)
		fmt.Println(x.Text('e', 10))
	}
}
