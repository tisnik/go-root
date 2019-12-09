// Seriál "Programovací jazyk Go"
//
// Dvacátá šestá část
//
// Demonstrační příklad číslo 21:
//    Porovnání dvou ukazatelů

package main

import "fmt"

func main() {
	var i1 interface{}

	fmt.Printf("%v %v\n", i1, i1 == nil)

	var i2 *int
	fmt.Printf("%v %v\n", i2, i2 == nil)

	fmt.Printf("%v\n", i1 == i2)
}
