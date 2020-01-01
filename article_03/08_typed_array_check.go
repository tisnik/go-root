// Seriál "Programovací jazyk Go"
//
// Třetí část
//
// Demonstrační příklad číslo 8:
//    Pole s prvky definovaného typu + kontrola typů.

package main

import "fmt"

// Mesic je nový datový typ se jménem viditelným mimo balíček
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

	var mesice2 [12]string = mesice

	fmt.Println(mesice)
	fmt.Println(mesice2)
}
