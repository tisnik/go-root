package main

import (
	"fmt"
	"math/big"
)

func main() {
	var x big.Rat
	var y big.Rat
	var z big.Rat
	var w big.Rat

	x.SetFrac64(1, 2)
	y.SetFrac64(1, 3)
	z.SetFrac64(2, 1)
	w.SetFrac64(4, 2)

	fmt.Println(x.String())
	fmt.Println(y.String())
	fmt.Println(z.String())
	fmt.Println(w.String())
}
