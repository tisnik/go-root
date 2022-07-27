// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Devadesátá třetí část
//    Knihovny s implementací generických datových typů pro programovací jazyk Go (2)
//    https://www.root.cz/clanky/knihovny-s-implementaci-generickych-datovych-typu-pro-programovaci-jazyk-go-2/
//
// Seznam příkladů ze devadesáté třetí části:
//    https://github.com/tisnik/go-root/blob/master/article_93/README.md

package main

import (
	"fmt"

	treeImpl "github.com/daichi-m/go18ds/trees/redblacktree"
)

func main() {
	tree := treeImpl.NewWithIntComparator[string]()
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
