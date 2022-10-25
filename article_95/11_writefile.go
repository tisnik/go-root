package main

import (
	"fmt"

	narray "github.com/akualab/narray/na64"
)

func main() {
	vector := narray.NewArray([]float64{1, 2, 3}, 3)
	matrix := narray.NewArray([]float64{1, 2, 3, 4, 5, 6}, 3, 2)
	cube := narray.NewArray([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 3, 2, 2)

	fmt.Println(vector)
	fmt.Println(matrix)
	fmt.Println(cube)

	vector.WriteFile("vector.dat")
	matrix.WriteFile("matrix.dat")
	cube.WriteFile("cube.dat")
}
