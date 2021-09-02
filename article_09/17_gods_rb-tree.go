// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Devátá část
//     Užitečné balíčky pro každodenní použití jazyka Go
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-jazyka-go/
//
// Demonstrační příklad číslo 17:
//    Datová struktura Red-Black tree.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_09/17_gods_rb-tree.html

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
