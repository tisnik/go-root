// Seriál "Programovací jazyk Go"
//
// Dvacátá šestá část
//
// Demonstrační příklad číslo 10B:
//    Ukazatel inicializovaný na nil

package main

import "fmt"

func main() {
	var v *int = nil

	fmt.Println(v)

	fmt.Println(*v)
}
