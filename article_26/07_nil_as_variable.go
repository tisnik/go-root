// Seriál "Programovací jazyk Go"
//
// Dvacátá šestá část
//
// Demonstrační příklad číslo 7:
//    nil může být korektní jméno konstanty.

package main

import "fmt"

func main() {
	fmt.Println(nil)

	nil := 42

	fmt.Println(nil)
}
