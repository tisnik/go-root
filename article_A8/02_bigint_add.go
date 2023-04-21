package main

import (
	"fmt"
	"math/big"
)

func main() {
	var x big.Int
	var y big.Int
	x.SetInt64(1)
	y.SetInt64(2)

	var z big.Int
	z.Add(&x, &y)

	fmt.Println(x.Text(10))
	fmt.Println(y.Text(10))
	fmt.Println(z.Text(10))
}
