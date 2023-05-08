package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	values := []float32{0, 1, -1, 10, -10, 100, 255, 256, 257, 2550, 2560, 2570, 1000, 1e38, -1e38}

	for _, value := range values {
		dec := decimal.NewFromFloat32(float32(value))
		bin, _ := dec.MarshalBinary()
		fmt.Println(dec, "\t", bin)
	}
}
