package main

import (
	"fmt"

	narray "github.com/akualab/narray/na64"
)

func main() {
	vector1 := narray.NewArray([]float64{1, 2, 3, 4, 5}, 5)

	sum := vector1.Sum()
	fmt.Println(sum)
}
