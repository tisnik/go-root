// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Seznam příkladů z devadesáté druhé části:
//    https://github.com/tisnik/go-root/blob/master/article_92/README.md

package main

import (
	"fmt"

	"github.com/daichi-m/go18ds/lists/arraylist"
)

func testContains(l *arraylist.List[string]) {
	fmt.Println("List:", l)
	for _, c := range []string{"a", "b", "c", "x"} {
		fmt.Printf("Contains '%s': %t\n", c, l.Contains(c))
	}
	fmt.Println()
}

func main() {
	l := arraylist.New[string]()
	testContains(l)

	l.Add("a")
	testContains(l)

	l.Add("b", "c")
	testContains(l)

	l.Clear()
	testContains(l)
}
