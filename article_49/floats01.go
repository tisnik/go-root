// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá devátá část
//    Popis vybraných balíčků nabízených projektem Gonum
//    https://www.root.cz/clanky/popis-vybranych-balicku-nabizenych-projektem-gonum/
//
// Demonstrační příklad číslo 1:
//

package main

import (
	"fmt"
	"gonum.org/v1/gonum/floats"
)

func main() {
	x := []float64{1, 2, 3, 4, 5, 6}

	fmt.Printf("x:    %v\n", x)
	fmt.Printf("Min:  %5.0f\n", floats.Min(x))
	fmt.Printf("Max:  %5.0f\n", floats.Max(x))
	fmt.Printf("Sum:  %5.0f\n", floats.Sum(x))
	fmt.Printf("Prod: %5.0f\n", floats.Prod(x))
}
