// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá devátá část
//    Popis vybraných balíčků nabízených projektem Gonum
//    https://www.root.cz/clanky/popis-vybranych-balicku-nabizenych-projektem-gonum/
//
// Seznam příkladů ze čtyřicáté deváté části:
//    https://github.com/tisnik/go-root/blob/master/article_49/README.md
//
// Demonstrační příklad číslo 7:
//

package main

import (
	"fmt"
	"gonum.org/v1/gonum/floats"
)

func main() {
	x := []float64{1, 2, 3, 4, 5, 6}
	y := []float64{10, 20, 30, 40, 50, 60}

	fmt.Printf("x:   %v\n", x)
	fmt.Printf("y:   %v\n\n", y)

	floats.Mul(x, y)
	fmt.Printf("x*y: %v\n", x)
}
