// Seriál "Programovací jazyk Go"
//
// Devátá část
//
// Demonstrační příklad číslo 14:
//    Datová struktura singlylinkedlist z knihovny GoDS.

package main

import (
	"fmt"
	sll "github.com/emirpasic/gods/lists/singlylinkedlist"
)

func printList(list *sll.List) {
	iterator := list.Iterator()
	for iterator.Next() {
		index, value := iterator.Index(), iterator.Value()
		fmt.Printf("item #%d == %s\n", index, value)
	}
	fmt.Println()
}

func main() {
	list := sll.New()
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
