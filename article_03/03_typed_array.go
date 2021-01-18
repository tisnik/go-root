// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třetí část
//    Datové typy v programovacím jazyku Go (2.část)
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go-2-cast/
//
// Demonstrační příklad číslo 3:
//    Pole s prvky definovaného typu.

package main

import "fmt"

// Mesic is type of item in array/slice
type Mesic string

func main() {
	var mesice [12]Mesic

	fmt.Println(mesice)

	mesice[0] = "Leden"
	mesice[1] = "Únor"
	mesice[2] = "Březen"
	mesice[3] = "Duben"
	mesice[4] = "Květen"
	mesice[5] = "Červen"
	mesice[6] = "Červenec"
	mesice[7] = "Srpen"
	mesice[8] = "Září"
	mesice[9] = "Říjen"
	mesice[10] = "Listopad"
	mesice[11] = "Prosinec"

	fmt.Println(mesice)
}
