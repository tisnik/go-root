// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá devátá část
//    Popis vybraných balíčků nabízených projektem Gonum
//    https://www.root.cz/clanky/popis-vybranych-balicku-nabizenych-projektem-gonum/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze čtyřicáté deváté části:
//    https://github.com/tisnik/go-root/blob/master/article_49/README.md

package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
)

func main() {
	values := []float64{1, 2, 3, 4, 5}
	fmt.Printf("Mean: %f\n", stat.Mean(values, nil))

	weights := []float64{1, 1, 1, 1, 1}
	fmt.Printf("Weighted mean #1: %f\n", stat.Mean(values, weights))

	weights2 := []float64{1, 1, 1, 1, 6}
	fmt.Printf("Weighted mean #2: %f\n", stat.Mean(values, weights2))
}
