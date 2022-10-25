package main

import (
	"fmt"

	narray "github.com/akualab/narray/na64"
)

func main() {
	vector1 := narray.NewArray([]float64{1, 2, 3, 4, 5}, 5)
	vector2 := narray.New(5)

	narray.AddConst(vector2, vector1, 100)
	fmt.Println(vector1)
	fmt.Println(vector2)
}
