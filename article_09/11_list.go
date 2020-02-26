// Seriál "Programovací jazyk Go"
//
// Devátá část
//     Užitečné balíčky pro každodenní použití jazyka Go
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-jazyka-go/
//
// Demonstrační příklad číslo 11:
//    Použití standardního balíčku "container/list"

package main

import (
	"container/list"
	"fmt"
)

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func main() {
	l := list.New()
	l.PushBack("foo")
	l.PushBack("bar")
	l.PushBack("baz")
	printList(l)
}
