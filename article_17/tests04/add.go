// Seriál "Programovací jazyk Go"
//
// Sedmnáctá část
//     Testování aplikací naprogramovaných v jazyce Go
//     https://www.root.cz/clanky/testovani-aplikaci-naprogramovanych-v-jazyce-go/
//
// Demonstrační příklad číslo 4:
//     Testovaný balíček.

package main

func add(x int, y int) int {
	return x - y
}

func main() {
	println(add(1, 2))
}
