// Seriál "Programovací jazyk Go"
//
// Třicátá čtvrtá část
//
// Demonstrační příklad číslo 19:
//    Deklarace konstant v Go.

package main

import "fmt"

const Pi float64 = 3.1415927
const E = 2.71828

const z0 int = 0
const z1 = 0

const z2 = z0 + z1

func main() {
	fmt.Printf("Pi = %f\n", Pi)
	fmt.Printf("e = %f\n", E)

	fmt.Printf("z0 = %d\n", z0)
	fmt.Printf("z1 = %d\n", z1)

	fmt.Printf("z2 = %d\n", z2)
}
