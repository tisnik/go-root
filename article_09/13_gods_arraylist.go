// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Devátá část
//     Užitečné balíčky pro každodenní použití jazyka Go
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-jazyka-go/
//
// Seznam příkladů z deváté části:
//    https://github.com/tisnik/go-root/blob/master/article_09/README.md
//
// Demonstrační příklad číslo 13:
//    Datová struktura arraylist z knihovny GoDS.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_09/13_gods_arraylist.html

package main

import (
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
)

func printList(list *arraylist.List) {
	iterator := list.Iterator()
	for iterator.Next() {
		index, value := iterator.Index(), iterator.Value()
		fmt.Printf("item #%d == %s\n", index, value)
	}
	fmt.Println()
}

func main() {
	list := arraylist.New()
	list.Add("a")
	list.Add("c", "b")

	printList(list)

	list.Swap(0, 1)
	list.Swap(1, 2)
	printList(list)

	list.Remove(2)
	printList(list)

	list.Remove(1)
	printList(list)
}
