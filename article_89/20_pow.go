// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Seznam příkladů z osmdesáté deváté části:
//    https://github.com/tisnik/go-root/blob/master/article_89/README.md

package main

import (
	"fmt"
	"math"
)

func pow(x float64, y float64) float64 {
	return math.Pow(x, y)
}

func main() {
	for x := 0.0; x < 5.0; x += 0.5 {
		fmt.Println(pow(x, 2))
	}
}
