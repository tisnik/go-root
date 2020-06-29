// Seriál "Programovací jazyk Go"
//
// Dvacátá šestá část
//    Problematika nulových hodnot v Go, aneb proč nil != nil
//    https://www.root.cz/clanky/problematika-nulovych-hodnot-v-go-aneb-proc-nil-nil/
//
// Demonstrační příklad číslo 9:
//    Hodnota nil a její význam (bez typu)

package main

import "fmt"

func main() {
	v := nil

	fmt.Println(v)
}
