// Seriál "Programovací jazyk Go"
//
// Dvacátá šestá část
//    Problematika nulových hodnot v Go, aneb proč nil != nil
//    https://www.root.cz/clanky/problematika-nulovych-hodnot-v-go-aneb-proc-nil-nil/
//
// Demonstrační příklad číslo 10B:
//    Ukazatel inicializovaný na nil

package main

import "fmt"

func main() {
	var v *int = nil

	fmt.Println(v)

	fmt.Println(*v)
}
