// Seriál "Programovací jazyk Go"
//
// Dvacátá šestá část
//    Problematika nulových hodnot v Go, aneb proč nil != nil
//    https://www.root.cz/clanky/problematika-nulovych-hodnot-v-go-aneb-proc-nil-nil/
//
// Demonstrační příklad číslo 7:
//    nil může být korektní jméno konstanty.

package main

import "fmt"

func main() {
	fmt.Println(nil)

	nil := 42

	fmt.Println(nil)
}
