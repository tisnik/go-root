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
	vector1 := narray.NewArray([]float64{1, 2, 3}, 3)
	vector2 := narray.NewArray([]float64{4, 5, 6}, 3)
	vector3 := narray.NewArray([]float64{7, 8, 9}, 3)

	dotProduct := narray.Dot(vector1, vector2, vector3)
	fmt.Println(dotProduct)
}
