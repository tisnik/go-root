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
// Demonstrační příklad číslo 25:
//    Čtení z nulového kanálu

package main

import "fmt"

func main() {
	var c1 chan int = nil

	fmt.Printf("%v %v\n", c1, c1 == nil)

	fmt.Printf("%d\n", <-c1)
}
