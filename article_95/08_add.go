package main

import (
	"fmt"

	narray "github.com/akualab/narray/na64"
)

func main() {
	vector1 := narray.NewArray([]float64{1, 2, 3, 4, 5}, 5)
	vector2 := narray.NewArray([]float64{1, 2, 3, 4, 5}, 5)
	vector3 := narray.New(5)

	narray.Add(vector3, vector1, vector2)
	fmt.Println(vector1)
	fmt.Println(vector2)
	fmt.Println(vector3)
}
