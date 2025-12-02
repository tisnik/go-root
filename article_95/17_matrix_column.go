// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Seznam příkladů z devadesáté páté části:
//    https://github.com/tisnik/go-root/blob/master/article_95/README.md

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

	col1 := trueMatrix.Col(nil, 0)
	fmt.Println(col1)

	col2 := trueMatrix.Col(nil, 1)
	fmt.Println(col2)
}
