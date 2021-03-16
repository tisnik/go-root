// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá šestá část
//    Problematika nulových hodnot v Go, aneb proč nil != nil
//    https://www.root.cz/clanky/problematika-nulovych-hodnot-v-go-aneb-proc-nil-nil/
//
// Demonstrační příklad číslo 26:
//    Nulové i normální kanály

package main

import "fmt"

func main() {
	var c1 chan int = nil
	var c2 chan int = make(chan int)

	fmt.Printf("%v %v\n", c1, c1 == nil)
	fmt.Printf("%v %v\n", c2, c2 == nil)
}
