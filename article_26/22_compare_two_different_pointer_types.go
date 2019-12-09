// Seriál "Programovací jazyk Go"
//
// Dvacátá šestá část
//
// Demonstrační příklad číslo 22:
//    Porovnání dvou ukazatelů různých typů

package main

import "fmt"

func main() {
	var i1 *int

	fmt.Printf("%v %v\n", i1, i1 == nil)

	var i2 *int32
	fmt.Printf("%v %v\n", i2, i2 == nil)

	fmt.Printf("%v\n", i1 == i2)
}
