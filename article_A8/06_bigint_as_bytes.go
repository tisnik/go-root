package main

import (
	"fmt"
	"math/big"
)

func main() {
	var x big.Int

	x.SetInt64(0)
	fmt.Println(x.Bytes())

	x.SetInt64(1)
	fmt.Println(x.Bytes())

	x.SetInt64(100)
	fmt.Println(x.Bytes())

	x.SetInt64(1000)
	fmt.Println(x.Bytes())

	x.SetInt64(100000000)
	fmt.Println(x.Bytes())

	x.SetInt64(100000000)
	x.Mul(&x, &x)
	fmt.Println(x.Bytes())
}
