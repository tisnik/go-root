package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewRat(1, 2)
	y := big.NewRat(1, 3)
	z := big.NewRat(2, 1)
	w := big.NewRat(4, 2)

	fmt.Println(x.String())
	fmt.Println(y.String())
	fmt.Println(z.String())
	fmt.Println(w.String())
}
