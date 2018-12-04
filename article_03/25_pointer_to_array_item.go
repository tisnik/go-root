// Seriál "Programovací jazyk Go"
//
// Třetí část
//
// Demonstrační příklad číslo 25:
//    Ukazatel na prvek pole.

package main

import "fmt"

func main() {
	mesice := [12]string{
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

	p_treti := &mesice[2]
	*p_treti = "March"

	fmt.Println(mesice)
}
