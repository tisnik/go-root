// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Devadesátá třetí část
//    Knihovny s implementací generických datových typů pro programovací jazyk Go (2)
//    https://www.root.cz/clanky/knihovny-s-implementaci-generickych-datovych-typu-pro-programovaci-jazyk-go-2/
//
// Seznam příkladů z devadesáté třetí části:
//    https://github.com/tisnik/go-root/blob/master/article_93/README.md

package main

import (
	"fmt"

	treeImpl "github.com/daichi-m/go18ds/trees/avltree"
)

func main() {
	tree := treeImpl.NewWithIntComparator[string]()

	tree.Put(1, "a")
	tree.Put(9, "i")
	tree.Put(2, "b")
	tree.Put(8, "h")
	tree.Put(3, "c")
	tree.Put(7, "g")
	tree.Put(4, "d")
	tree.Put(6, "f")
	tree.Put(5, "e")

	fmt.Println(tree)

	it := tree.Iterator()
	for it.Next() {
		value := it.Value()
		fmt.Printf("%3s \t %T\n", value, value)
	}
}
