package main

import (
	"fmt"

	narray "github.com/akualab/narray/na64"
)

func main() {
	matrix := narray.NewArray([]float64{1, 2, 3, 4, 5, 6}, 3, 2)

	trueMatrix := matrix.Matrix(-1, -1)
	fmt.Println(matrix)
	fmt.Println(trueMatrix)
}
