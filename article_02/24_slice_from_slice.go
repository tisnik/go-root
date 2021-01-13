// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Druhá část:
//    Datové typy v programovacím jazyku Go
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 24
//    Kopie řezů polí

package main

import "fmt"

func main() {
	var a [10]int

	slice1 := a[4:9]
	slice2 := slice1[3:]

	fmt.Printf("Pole:            %v\n", a)
	fmt.Printf("Delka pole:      %d\n\n", len(a))

	fmt.Printf("Rez 1:           %v\n", slice1)
	fmt.Printf("Delka rezu 1:    %d\n", len(slice1))
	fmt.Printf("Kapacita rezu 1: %d\n\n", cap(slice1))

	fmt.Printf("Rez 2:           %v\n", slice2)
	fmt.Printf("Delka rezu 2:    %d\n", len(slice2))
	fmt.Printf("Kapacita rezu 2: %d\n\n", cap(slice2))

	slice2[0] = 99
	slice2[1] = 99

	fmt.Printf("Pole:            %v\n", a)
	fmt.Printf("Rez 1:           %v\n", slice1)
	fmt.Printf("Rez 2:           %v\n", slice2)
}
