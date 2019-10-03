// Seriál "Programovací jazyk Go"
//
// Dvacátá šestá část
//
// Demonstrační příklad číslo 10:
//    Ukazatel inicializovaný na nil

package main

import "fmt"

func main() {
	var v *int = nil

	fmt.Println(v)

	*v = 42
}
