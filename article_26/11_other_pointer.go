// Seriál "Programovací jazyk Go"
//
// Dvacátá šestá část
//
// Demonstrační příklad číslo 11:
//    Ukazatel na celé číslo

package main

import "fmt"

func main() {
	var v *int = 1234

	fmt.Println(v)

	*v = 42
}
