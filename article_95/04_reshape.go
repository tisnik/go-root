package main

import (
	"fmt"

	narray "github.com/akualab/narray/na64"
)

func main() {
	vector := narray.NewArray([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 12)
	matrix := vector.Reshape(4, 3)

	fmt.Println(vector)
	fmt.Println(matrix)
}
