package main

import (
	"fmt"
	"math/big"
)

func main() {
	var x big.Int
	x.SetInt64(2)

	var z big.Int
	z.SetInt64(1)

	for i := 0; i < 100; i++ {
		z.Mul(&z, &x)
		fmt.Println(z.Text(10))
	}
}
