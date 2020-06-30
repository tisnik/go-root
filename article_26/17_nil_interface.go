// Seriál "Programovací jazyk Go"
//
// Dvacátá šestá část
//    Problematika nulových hodnot v Go, aneb proč nil != nil
//    https://www.root.cz/clanky/problematika-nulovych-hodnot-v-go-aneb-proc-nil-nil/
//
// Demonstrační příklad číslo 17:
//    Prázdné rozhraní

package main

import "fmt"

func main() {
	var i1 interface{}

	fmt.Printf("%v %v\n", i1, i1 == nil)
}
