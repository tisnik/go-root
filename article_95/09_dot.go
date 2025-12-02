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
	vector1 := narray.NewArray([]float64{1, 2, 3, 4, 5}, 5)

	dotProduct := narray.Dot(vector1, vector1)
	fmt.Println(dotProduct)
	fmt.Println(1*1 + 2*2 + 3*3 + 4*4 + 5*5)
}
