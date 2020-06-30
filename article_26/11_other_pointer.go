// Seriál "Programovací jazyk Go"
//
// Dvacátá šestá část
//    Problematika nulových hodnot v Go, aneb proč nil != nil
//    https://www.root.cz/clanky/problematika-nulovych-hodnot-v-go-aneb-proc-nil-nil/
//
// Demonstrační příklad číslo 11:
//    Ukazatel na celé číslo

package main

import "fmt"

func main() {
	var v *int = 1234

	fmt.Println(v)

	*v = 42
}
