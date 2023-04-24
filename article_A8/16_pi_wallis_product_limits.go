package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	var result big.Rat

	result.SetFrac64(2, 1)

	for n := int64(1); n < math.MaxInt64; n++ {
		m := 4 * n * n

		var item big.Rat
		item.SetFrac64(m, m-1)
		result.Mul(&result, &item)

		f, _ := result.Float64()
		absError := math.Pi - f
		relError := 100.0 * absError / math.Pi
		fmt.Println(f, "\t", absError, "\t", relError)
	}
}
