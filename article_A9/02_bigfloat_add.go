package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewFloat(1)
	y := big.NewFloat(0.1)

	for i := 0; i < 20; i++ {
		x.Add(x, y)
		fmt.Println(x.Text('f', 10))
	}
}
