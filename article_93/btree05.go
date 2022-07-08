// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/

package main

import (
	"fmt"

	treeImpl "github.com/daichi-m/go18ds/trees/btree"
)

func main() {
	tree := treeImpl.NewWithIntComparator[string](3)

	tree.Put(1, "a")
	tree.Put(9, "i")
	tree.Put(2, "b")
	tree.Put(8, "h")
	tree.Put(3, "c")
	tree.Put(7, "g")
	tree.Put(4, "d")
	tree.Put(6, "f")
	tree.Put(5, "e")

	fmt.Println(tree.Values())
}
