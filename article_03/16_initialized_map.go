// Seriál "Programovací jazyk Go"
//
// Třetí část
//
// Demonstrační příklad číslo 16:
//    Mapa s inicializací.

package main

import "fmt"

func main() {
	var m1 map[int]string = make(map[int]string)
	fmt.Println(m1)

	m1[0] = "nula"
	m1[1] = "jedna"
	m1[2] = "dva"
	m1[3] = "tri"
	m1[4] = "ctyri"
	m1[5] = "pet"
	m1[6] = "sest"

	fmt.Println(m1)
}
