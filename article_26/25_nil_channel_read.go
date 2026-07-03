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
// Demonstrační příklad číslo 25:
//    Čtení z nulového kanálu.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_26/25_nil_channel_read.html

package main

import "fmt"

func main() {
	var c1 chan int = nil

	fmt.Printf("%v %v\n", c1, c1 == nil)

	fmt.Printf("%d\n", <-c1)
}
