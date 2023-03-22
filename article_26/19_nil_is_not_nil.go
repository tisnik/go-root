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
// Demonstrační příklad číslo 19:
//    nil se někdy nerovná nil

package main

import "fmt"

type T *int

func main() {
	var i1 interface{}

	fmt.Printf("%v %v\n", i1, i1 == nil)

	var i2 T = nil
	fmt.Printf("%v %v\n", i2, i2 == nil)

	fmt.Printf("%v\n", i1 == i2)
}
