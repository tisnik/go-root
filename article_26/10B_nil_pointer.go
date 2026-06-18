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
// Demonstrační příklad číslo 10B:
//    Ukazatel inicializovaný na nil.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_26/10B_nil_pointer.html

package main

import "fmt"

func main() {
	var v *int = nil

	fmt.Println(v)

	fmt.Println(*v)
}
