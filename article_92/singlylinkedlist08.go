// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Seznam příkladů z devadesáté druhé části:
//    https://github.com/tisnik/go-root/blob/master/article_92/README.md

package main

import (
	"fmt"

	"github.com/daichi-m/go18ds/lists/singlylinkedlist"
)

func main() {
	l := singlylinkedlist.New[string]()
	fmt.Println(l)

	l.Add("a")
	fmt.Println(l)

	l.Add("b", "c")
	fmt.Println(l)

	for i := 0; i < 5; i++ {
		l.Remove(0)
		fmt.Println(l)
	}
}
