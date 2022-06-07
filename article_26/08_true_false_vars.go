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
// Demonstrační příklad číslo 8:
//    true a false mohou být korektní jméno konstanty.

package main

import "fmt"

func main() {
	fmt.Println(true)
	fmt.Println(false)

	x := true
	true := false
	false := x

	fmt.Println("oh no...")

	fmt.Println(true)
	fmt.Println(false)
}
