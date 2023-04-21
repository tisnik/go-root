package main

import (
	"fmt"
	"math/big"
)

func main() {
	var x big.Int
	x.SetInt64(100000000000)
	x.Mul(&x, &x)

	fmt.Println(x.Bits())

	words := x.Bits()
	words[0] = 1
	words[1] = 0

	fmt.Println(x.Bits())
	fmt.Println(x.Text(10))

	words[0] = 0
	words[1] = 1

	fmt.Println(x.Bits())
	fmt.Println(x.Text(10))
}
