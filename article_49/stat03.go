// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá devátá část
//    Popis vybraných balíčků nabízených projektem Gonum
//    https://www.root.cz/clanky/popis-vybranych-balicku-nabizenych-projektem-gonum/

package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
)

func main() {
	values := []float64{1, 2, 3, 4, 5}
	fmt.Printf("Harmonic mean: %f\n", stat.HarmonicMean(values, nil))

	weights := []float64{1, 1, 1, 1, 1}
	fmt.Printf("Weighted Harmonic mean #1: %f\n", stat.HarmonicMean(values, weights))

	weights2 := []float64{1, 1, 1, 1, 6}
	fmt.Printf("Weighted Harmonic mean #2: %f\n", stat.HarmonicMean(values, weights2))
}
