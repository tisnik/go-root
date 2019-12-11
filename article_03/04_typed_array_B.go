// Seriál "Programovací jazyk Go"
//
// Třetí část
//
// Demonstrační příklad číslo 4:
//    Pole s prvky definovaného typu.

package main

import "fmt"

// Mesic is type of item in array/slice
type Mesic string

func main() {
	mesice := [12]Mesic{
		"Leden",
		"Únor",
		"Březen",
		"Duben",
		"Květen",
		"Červen",
		"Červenec",
		"Srpen",
		"Září",
		"Říjen",
		"Listopad",
		"Prosinec"}

	fmt.Println(mesice)
}
