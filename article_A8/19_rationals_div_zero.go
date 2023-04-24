package main

import (
	"fmt"
	"math/big"
)

func main() {
	var x big.Rat

	x.SetFrac64(0, 1)

	fmt.Println(x.String())

	x.Inv(&x)
	fmt.Println(x.String())
}
