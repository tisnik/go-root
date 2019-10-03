// Seriál "Programovací jazyk Go"
//
// Dvacátá šestá část
//
// Demonstrační příklad číslo 16:
//    Prázdná mapa a přidání prvků do prázdné mapy

package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	fmt.Printf("%v %v\n", m1, m1 == nil)

	m1["foo"] = 3
	fmt.Printf("%v %v\n", m1, m1 == nil)
}
