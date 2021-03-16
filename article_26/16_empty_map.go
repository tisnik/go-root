// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá šestá část
//    Problematika nulových hodnot v Go, aneb proč nil != nil
//    https://www.root.cz/clanky/problematika-nulovych-hodnot-v-go-aneb-proc-nil-nil/
//
// Demonstrační příklad číslo 16:
//    Prázdná mapa a přidání prvků do prázdné mapy

package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	fmt.Printf("%v %v\n", m1, m1 == nil)

	m1["foo"] = 3
	fmt.Printf("%v %v\n", m1, m1 == nil)
}
