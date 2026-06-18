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
// Demonstrační příklad číslo 7:
//    nil může být korektní jméno konstanty.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_26/07_nil_as_variable.html

package main

import "fmt"

func main() {
	fmt.Println(nil)

	nil := 42

	fmt.Println(nil)
}
