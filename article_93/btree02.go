package main

import (
	"fmt"

	treeImpl "github.com/daichi-m/go18ds/trees/btree"
)

func main() {
	tree := treeImpl.NewWithIntComparator[string](4)
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
