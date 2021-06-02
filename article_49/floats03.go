// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá devátá část
//    Popis vybraných balíčků nabízených projektem Gonum
//    https://www.root.cz/clanky/popis-vybranych-balicku-nabizenych-projektem-gonum/
//
// Demonstrační příklad číslo 3:
//

package main

import (
	"fmt"
	"gonum.org/v1/gonum/floats"
)

func main() {
	x := []float64{1, 2, 3, 4, 5, 6}

	fmt.Printf("x:        %v\n", x)

	floats.Reverse(x)
	fmt.Printf("reversed: %v\n", x)

	floats.Scale(10, x)
	fmt.Printf("scaled:   %v\n", x)
}
