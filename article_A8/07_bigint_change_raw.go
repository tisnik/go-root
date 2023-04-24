package main

import (
	"fmt"
	"math/big"
)

func main() {
	var x big.Int
	x.SetInt64(100000000000)
	x.Mul(&x, &x)

	fmt.Println(x.Text(10))
	fmt.Println(x.Bits())
	fmt.Println()

	words := x.Bits()
	words[0] = 1
	words[1] = 0

	fmt.Println(x.Text(10))
	fmt.Println(x.Bits())
	fmt.Println()

	words[0] = 0
	words[1] = 1

	fmt.Println(x.Text(10))
	fmt.Println(x.Bits())
	fmt.Println()

	words[0] = 1000
	words[1] = 1000

	fmt.Println(x.Text(10))
	fmt.Println(x.Bits())
}
