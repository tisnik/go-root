// Seriál "Programovací jazyk Go"
//
// Druhá část:
//    Datové typy v programovacím jazyku Go
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 20
//    Základní operace s poli

package main

import "fmt"

func main() {
	var a1 [10]byte
	var a2 [10]int32
	a3 := [10]int32{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}

	fmt.Printf("Delka pole 1: %d\n", len(a1))
	fmt.Printf("Delka pole 2: %d\n", len(a2))
	fmt.Printf("Delka pole 3: %d\n", len(a3))

	var a [10]int

	fmt.Printf("Pole pred upravou: %v\n", a)

	for i := 0; i < len(a1); i++ {
		a[i] = i * 2
	}

	fmt.Printf("Pole po uprave:    %v\n", a)

	var matice [10][10]float32
	fmt.Printf("Matice:    %v\n", matice)
}
