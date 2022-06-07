// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá šestá část
//    Problematika nulových hodnot v Go, aneb proč nil != nil
//    https://www.root.cz/clanky/problematika-nulovych-hodnot-v-go-aneb-proc-nil-nil/
//
// Seznam příkladů z dvacáté šesté části:
//    https://github.com/tisnik/go-root/blob/master/article_26/README.md
//
// Demonstrační příklad číslo 13:
//    Nulová mapa

package main

import "fmt"

func main() {
	var m1 map[string]int = nil
	var m2 map[string]int

	fmt.Printf("%v\n", m1)
	fmt.Printf("%v\n", m2)
}
