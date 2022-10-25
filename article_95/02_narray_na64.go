package main

import (
	"fmt"

	narray "github.com/akualab/narray/na64"
)

func main() {
	scalar := narray.New()
	vector := narray.New(10)
	matrix := narray.New(4, 3)
	cube := narray.New(3, 4, 5)

	fmt.Println(scalar)
	fmt.Println(vector)
	fmt.Println(matrix)
	fmt.Println(cube)
}
