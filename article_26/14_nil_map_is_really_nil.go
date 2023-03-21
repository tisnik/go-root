// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá šestá část
//    Problematika nulových hodnot v Go, aneb proč nil != nil
//    https://www.root.cz/clanky/problematika-nulovych-hodnot-v-go-aneb-proc-nil-nil/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z dvacáté šesté části:
//    https://github.com/tisnik/go-root/blob/master/article_26/README.md
//
// Demonstrační příklad číslo 14:
//    Nulová mapa

package main

import "fmt"

func main() {
	var m1 map[string]int = nil
	var m2 map[string]int

	fmt.Printf("%v %v\n", m1, m1 == nil)
	fmt.Printf("%v %v\n", m2, m2 == nil)
}
