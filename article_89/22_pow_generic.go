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

type floats interface {
	float32 | float64
}

func pow[T floats](x T, y T) T {
	return T(math.Pow(float64(x), float64(y)))
}

func main() {
	for x := float32(0.0); x < 5.0; x += 0.5 {
		fmt.Println(pow(x, 2))
	}

	fmt.Println()

	for x := 0.0; x < 5.0; x += 0.5 {
		fmt.Println(pow(x, 2))
	}
}
