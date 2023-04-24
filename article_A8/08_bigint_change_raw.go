package main

import (
	"fmt"
	"math/big"
)

func main() {
	var x big.Int
	x.SetInt64(0)

	fmt.Println(x.Bytes())
	fmt.Println(x.Text(10))
	fmt.Println()

	x.SetBytes([]byte{})
	fmt.Println(x.Bytes())
	fmt.Println(x.Text(10))
	fmt.Println()

	x.SetBytes([]byte{1, 0})
	fmt.Println(x.Bytes())
	fmt.Println(x.Text(10))
	fmt.Println()

	x.SetBytes([]byte{1, 0, 0})
	fmt.Println(x.Bytes())
	fmt.Println(x.Text(10))
	fmt.Println()

	x.SetBytes([]byte{1, 0, 0, 0})
	fmt.Println(x.Bytes())
	fmt.Println(x.Text(10))
	fmt.Println()

	x.SetBytes([]byte{1, 0, 0, 0, 0})
	fmt.Println(x.Bytes())
	fmt.Println(x.Text(10))
	fmt.Println()

	x.SetBytes([]byte{1, 0, 0, 0, 0, 0})
	fmt.Println(x.Bytes())
	fmt.Println(x.Text(10))
	fmt.Println()

	x.SetBytes([]byte{1, 0, 0, 0, 0, 0, 0})
	fmt.Println(x.Bytes())
	fmt.Println(x.Text(10))
	fmt.Println()

	x.SetBytes([]byte{1, 0, 0, 0, 0, 0, 0, 0})
	fmt.Println(x.Bytes())
	fmt.Println(x.Text(10))
	fmt.Println()

	x.SetBytes([]byte{1, 0, 0, 0, 0, 0, 0, 0, 0})
	fmt.Println(x.Bytes())
	fmt.Println(x.Text(10))
	fmt.Println()
}
