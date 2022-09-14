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

func testContains(l *singlylinkedlist.List[string]) {
	fmt.Println("List:", l)
	for _, c := range []string{"a", "b", "c", "x"} {
		fmt.Printf("Contains '%s': %t\n", c, l.Contains(c))
	}
	fmt.Println()
}

func main() {
	l := singlylinkedlist.New[string]()
	testContains(l)

	l.Add("a")
	testContains(l)

	l.Add("b", "c")
	testContains(l)

	l.Clear()
	testContains(l)
}
