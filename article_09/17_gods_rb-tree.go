// Seriál "Programovací jazyk Go"
//
// Devátá část
//
// Demonstrační příklad číslo 17:
//    Datová struktura Red-Black tree.

package main

import (
	"fmt"
	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

func main() {
	tree := rbt.NewWithIntComparator()
	fmt.Println(tree)

	tree.Put(1, "G")
	fmt.Println(tree)

	tree.Put(2, "a")
	tree.Put(3, "b")
	fmt.Println(tree)

	tree.Put(4, "a")
	tree.Put(5, "b")
	fmt.Println(tree)

	tree.Put(6, "a")
	tree.Put(7, "b")
	fmt.Println(tree)

	tree.Put(8, "a")
	tree.Put(9, "b")
	fmt.Println(tree)
}
