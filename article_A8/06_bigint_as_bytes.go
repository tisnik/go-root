package main

import (
	"fmt"
	"math/big"
)

func main() {
	var x big.Int
	x.SetInt64(100000000)

	fmt.Println(x.Bytes())
}
