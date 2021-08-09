// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třetí část
//    Datové typy v programovacím jazyku Go (2.část)
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go-2-cast/
//
// Demonstrační příklad číslo 7:
//    Pole s prvky definovaného typu + kontrola typů.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_03/07_typed_array_check.html

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
