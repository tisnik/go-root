// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Druhá část:
//    Datové typy v programovacím jazyku Go
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 22
//    Řezy polí

package main

import "fmt"

func main() {
	var a1 [100]byte
	var a2 [100]int32

	fmt.Printf("Delka pole 1: %d\n", len(a1))
	fmt.Printf("Delka pole 2: %d\n", len(a2))

	var slice1 []byte = a1[10:20]
	fmt.Printf("Delka rezu 1:    %d\n", len(slice1))
	fmt.Printf("Kapacita rezu 1: %d\n", cap(slice1))

	var slice2 = a1[20:30]
	fmt.Printf("Delka rezu 2:    %d\n", len(slice2))
	fmt.Printf("Kapacita rezu 2: %d\n", cap(slice2))

	slice3 := a1[30:40]
	fmt.Printf("Delka rezu 3:    %d\n", len(slice3))
	fmt.Printf("Kapacita rezu 3: %d\n", cap(slice3))
}
