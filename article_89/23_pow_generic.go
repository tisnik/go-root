// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/

package main

import (
	"fmt"
	"math"
)

type floats interface {
	float32 | float64
}

type numeric interface {
	float32 | float64 | int
}

func pow[T floats, U numeric](x T, y U) T {
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
