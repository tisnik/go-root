package main

import (
	"fmt"

	narray "github.com/akualab/narray/na64"
)

func main() {
	vector1 := narray.NewArray([]float64{1, 2, 3, 4, 5}, 5)
	vector2 := narray.NewArray([]float64{1, 2, 3, 4, 5}, 5)

	narray.AddScaled(vector1, vector2, 100)
	fmt.Println(vector1)
	fmt.Println(vector2)
}
