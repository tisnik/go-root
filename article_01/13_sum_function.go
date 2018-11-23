// Seriál "Programovací jazyk Go"
//
// První část
//
// Demonstrační příklad číslo 13:
//    Funkce se dvěma parametry

package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("sum =", sum(1, 2))
}
