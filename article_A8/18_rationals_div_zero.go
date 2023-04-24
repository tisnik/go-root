package main

import (
	"fmt"
	"math/big"
)

func main() {
	var x big.Rat

	x.SetFrac64(1, 0)

	fmt.Println(x.String())
}
