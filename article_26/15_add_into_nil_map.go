// Seriál "Programovací jazyk Go"
//
// Dvacátá šestá část
//
// Demonstrační příklad číslo 15:
//    Přidání prvků do nulové mapy

package main

import "fmt"

func main() {
	var m1 map[string]int = nil
	fmt.Printf("%v %v\n", m1, m1 == nil)

	m1["foo"] = 3
}
