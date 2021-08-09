// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třetí část
//    Datové typy v programovacím jazyku Go (2.část)
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go-2-cast/
//
// Demonstrační příklad číslo 25:
//    Ukazatel na prvek pole.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_03/25_pointer_to_array_item.html

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
