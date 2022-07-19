// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Seznam příkladů z devadesáté druhé části:
//    https://github.com/tisnik/go-root/blob/master/article_92/README.md

package main

import (
	"fmt"

	"github.com/daichi-m/go18ds/lists/doublylinkedlist"
)

func main() {
	l := doublylinkedlist.New[string]("zoo", "aleph", "foo", "bar", "baz")

	it := l.Iterator()
	for it.Next() {
		value := it.Value()
		fmt.Printf("%3s \t %T\n", value, value)
	}

}
