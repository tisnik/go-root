package main

import (
	"fmt"
	"math/big"
)

func main() {
	var x big.Rat
	var y big.Rat
	var z big.Rat

	x.SetFrac64(1, 2)
	y.SetFrac64(1, 1)
	z.SetFrac64(1, 1)

	for i := 0; i < 10; i++ {
		z.Mul(&z, &x)
		z.Add(&z, &y)

		fmt.Println(z.Float64())
	}

}
