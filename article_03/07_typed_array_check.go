// Seriál "Programovací jazyk Go"
//
// Třetí část
//    Datové typy v programovacím jazyku Go (2.část)
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go-2-cast/
//
// Demonstrační příklad číslo 7:
//    Pole s prvky definovaného typu + kontrola typů.

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

	var mesice2 [12]string = mesice

	fmt.Println(mesice)
	fmt.Println(mesice2)
}
