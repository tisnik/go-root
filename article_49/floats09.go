// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá devátá část
//    Popis vybraných balíčků nabízených projektem Gonum
//    https://www.root.cz/clanky/popis-vybranych-balicku-nabizenych-projektem-gonum/
//
// Demonstrační příklad číslo 9:
//

package main

import (
	"fmt"
	"gonum.org/v1/gonum/floats"
)

func main() {
	x := []float64{1, 2, 3, 4, 5, 6}
	y := []float64{1, 2, 3, 4, 5, 6}
	z := []float64{1.0, 2.1, 2.9, 4.1, 4.9, 6.0}

	fmt.Printf("x:   %v\n", x)
	fmt.Printf("y:   %v\n", y)
	fmt.Printf("z:   %v\n", z)

	fmt.Printf("x==y?: %t\n", floats.Equal(x, y))
	fmt.Printf("x==z?: %t\n", floats.Equal(x, z))
}
