// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Devátá část
//     Užitečné balíčky pro každodenní použití jazyka Go
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-jazyka-go/
//
// Demonstrační příklad číslo 11:
//    Použití standardního balíčku "container/list"
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_09/11_list.html

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
