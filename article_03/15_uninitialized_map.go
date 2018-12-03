// Seriál "Programovací jazyk Go"
//
// Třetí část
//
// Demonstrační příklad číslo 15:
//    Mapa bez inicializace.

package main

import "fmt"

func main() {
	var m1 map[int]string
	fmt.Println(m1)

	m1[0] = "nula"
}
