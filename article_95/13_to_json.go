package main

import (
	"fmt"
	"log"

	narray "github.com/akualab/narray/na64"
)

func main() {
	vector := narray.NewArray([]float64{1, 2, 3}, 3)
	matrix := narray.NewArray([]float64{1, 2, 3, 4, 5, 6}, 3, 2)
	cube := narray.NewArray([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 3, 2, 2)

	j, err := vector.ToJSON()
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(j)

	j, err = matrix.ToJSON()
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(j)

	j, err = cube.ToJSON()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(j)
}
