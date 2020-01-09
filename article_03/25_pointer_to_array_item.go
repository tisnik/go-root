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

	pThirdMonth := &mesice[2]
	*pThirdMonth = "March"

	fmt.Println(mesice)
}
