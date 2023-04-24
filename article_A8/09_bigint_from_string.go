package main

import (
	"fmt"
	"math/big"
)

func main() {
	var x big.Int

	x.SetString("42", 10)
	fmt.Println(x.Text(10))

	x.SetString("1000000000000000000000000000", 10)
	fmt.Println(x.Text(10))

	x.SetString("1000000000000000000000000000", 2)
	fmt.Println(x.Text(10))

	x.SetString("0xcafebabe", 0)
	fmt.Println(x.Text(10))
}
