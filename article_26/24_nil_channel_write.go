// Seriál "Programovací jazyk Go"
//
// Dvacátá šestá část
//    Problematika nulových hodnot v Go, aneb proč nil != nil
//    https://www.root.cz/clanky/problematika-nulovych-hodnot-v-go-aneb-proc-nil-nil/
//
// Demonstrační příklad číslo 24:
//    Zápis do nulového kanálu

package main

import "fmt"

func main() {
	var c1 chan int = nil

	fmt.Printf("%v %v\n", c1, c1 == nil)

	c1 <- 10
}
