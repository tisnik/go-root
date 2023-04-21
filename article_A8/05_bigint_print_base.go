package main

import (
	"fmt"
	"math/big"
)

func main() {
	var x big.Int
	x.SetInt64(10000000000)

	for base := 2; base <= 36; base++ {
		fmt.Println(base, x.Text(base))
	}
}
