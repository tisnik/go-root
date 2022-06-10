// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá čtvrtá část
//    Programovací jazyk Go pro skalní céčkaře
//    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare/
//
// Seznam příkladů ze třicáté čtvrté části:
//    https://github.com/tisnik/go-root/blob/master/article_34/README.md
//
// Demonstrační příklad číslo 20:
//    Deklarace konstant v Go.

package main

import "fmt"

const (
	// Pi is the classic math constant Pi
	Pi float64 = 3.1415927
	E          = 2.71828

	z0 int = 0
	z1     = 0
	z2     = z0 + z1
)

func main() {
	fmt.Printf("Pi = %f\n", Pi)
	fmt.Printf("e = %f\n", E)

	fmt.Printf("z0 = %d\n", z0)
	fmt.Printf("z1 = %d\n", z1)

	fmt.Printf("z2 = %d\n", z2)
}
