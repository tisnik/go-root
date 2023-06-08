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
	xs := []float64{1, 2, 3, 4, 5}
	ys := []float64{2, 3, 4, 5, 6}
	offset, slope := stat.LinearRegression(xs, ys, nil, false)
	fmt.Printf("Offset:  %f\n", offset)
	fmt.Printf("Slope:   %f\n", slope)
}
