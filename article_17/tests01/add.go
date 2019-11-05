// Seriál "Programovací jazyk Go"
//
// Sedmnáctá část
//
// Demonstrační příklad číslo 1:
//     Testovaný balíček.

package main

func Add(x int, y int) int {
	return x + y
}

func main() {
	println(Add(1, 2))
}
