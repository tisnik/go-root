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
	y.SetFrac64(1, 3)
	z.Add(&x, &y)

	fmt.Println(x.String())
	fmt.Println(y.String())
	fmt.Println(z.String())
}
