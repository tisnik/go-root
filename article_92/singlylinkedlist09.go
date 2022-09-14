// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Devadesátá druhá část
//    Knihovny s implementací generických datových typů pro programovací jazyk Go (2)
//    https://www.root.cz/clanky/knihovny-s-implementaci-generickych-datovych-typu-pro-programovaci-jazyk-go-2/
//
// Seznam příkladů z devadesáté druhé části:
//    https://github.com/tisnik/go-root/blob/master/article_92/README.md

package main

import (
	"fmt"

	"github.com/daichi-m/go18ds/lists/singlylinkedlist"
)

func main() {
	l := singlylinkedlist.New[string]("foo", "bar", "baz")
	fmt.Println(l)
	fmt.Println()

	l.Swap(0, 1)
	fmt.Println(l)
	fmt.Println()

	l.Swap(1, 2)
	fmt.Println(l)
	fmt.Println()

	l.Swap(2, 3)
	fmt.Println(l)
	fmt.Println()
}
