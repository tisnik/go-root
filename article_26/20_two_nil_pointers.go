// Seriál "Programovací jazyk Go"
//
// Dvacátá šestá část
//    Problematika nulových hodnot v Go, aneb proč nil != nil
//    https://www.root.cz/clanky/problematika-nulovych-hodnot-v-go-aneb-proc-nil-nil/
//
// Demonstrační příklad číslo 20:
//    Dva ukazatele s nulovou hodnotou

package main

import "fmt"

func main() {
	var p1 *int
	var p2 *int

	fmt.Printf("%v %v\n", p1, p1 == nil)
	fmt.Printf("%v %v\n", p2, p2 == nil)
	fmt.Printf("%v\n", p1 == p2)
}
