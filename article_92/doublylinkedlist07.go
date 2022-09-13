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

	"github.com/daichi-m/go18ds/lists/doublylinkedlist"
)

func testGet(l *doublylinkedlist.List[string]) {
	fmt.Println("List:", l)

	for i := -1; i < l.Size()+1; i++ {
		value, found := l.Get(i)
		if found {
			fmt.Printf("Item %d = '%s'\n", i, value)
		} else {
			fmt.Printf("Item %d not found\n", i)
		}
	}
	fmt.Println()
}

func main() {
	l := doublylinkedlist.New[string]()
	testGet(l)

	l.Add("a")
	testGet(l)

	l.Add("b", "c")
	testGet(l)

	l.Clear()
	testGet(l)
}
