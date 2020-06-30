// Seriál "Programovací jazyk Go"
//
// Dvacátá šestá část
//    Problematika nulových hodnot v Go, aneb proč nil != nil
//    https://www.root.cz/clanky/problematika-nulovych-hodnot-v-go-aneb-proc-nil-nil/
//
// Demonstrační příklad číslo 15:
//    Přidání prvků do nulové mapy

package main

import "fmt"

func main() {
	var m1 map[string]int = nil
	fmt.Printf("%v %v\n", m1, m1 == nil)

	m1["foo"] = 3
}
